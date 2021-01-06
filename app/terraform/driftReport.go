package terraform

import (
	"fmt"
	"os"
	"regexp"
	"strconv"

	tfjson "github.com/hashicorp/terraform-json"
	tail "github.com/hpcloud/tail"

	"clairvoyance/log"
)

// Parse out.tfplan and return the last line if it contains "Plan".
func GetResourceModificationCount(planFileRawString string) string {
	var filename string = "/tmp/clairvoyance-tmp"
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	} else {
		file.WriteString(planFileRawString)
	}
	file.Close()

	t, err := tail.TailFile(filename, tail.Config{Follow: true})
	for line := range t.Lines {
		matched, err := regexp.MatchString("\\bPlan\\b", line.Text)
		if err != nil {
			panic(err)
		} else if matched {
			return line.Text
		}
	}
	return ""
}

// Get # of Add, Change, Destroy in out.tfplan
func ParseResourceModificationCount(resourceModificationString string) map[string]int {
	var resourceModification map[string]int
	resourceModification = make(map[string]int)

	re := regexp.MustCompile("[0-9]+")
	var counts []string = re.FindAllString(resourceModificationString, -1)

	countAdd, err := strconv.Atoi(counts[0])
	if err != nil {
		panic(err)
	}
	countChange, err := strconv.Atoi(counts[1])
	if err != nil {
		panic(err)
	}
	countDestroy, err := strconv.Atoi(counts[2])
	if err != nil {
		panic(err)
	}

	resourceModification["CountAdd"] = countAdd
	resourceModification["CountChange"] = countChange
	resourceModification["CountDestroy"] = countDestroy

	return resourceModification
}

// terraform plan -detailed-exitcode (essentially)
// true == diff || false == No changes.
func GetDriftSummary(exitStatus bool, state *tfjson.State) string {
	var message string
	if exitStatus {
		message = "Drift detected for Plan."
		log.Debugf("[GetDriftSummary] %s", message)
	} else if !exitStatus {
		message = "No changes."
		log.Debugf("[GetDriftSummary] %s", message)
	} else {
		message = "Error planning Terraform project."
		log.Debugf("[GetDriftSummary] %s", message)
	}
	return message
}

// Populate a TerraformService structure with relevant data.
func UpdateDriftReportData(state *tfjson.State, projectName string, counts map[string]int, summary string) *TerraformService {
	tfs := &TerraformService{
		//State:            state,
		ProjectName:      projectName,
		TerraformVersion: state.TerraformVersion,
		CountAdd:         counts["CountAdd"],
		CountChange:      counts["CountChange"],
		CountDestroy:     counts["CountDestroy"],
		Summary:          summary,
	}
	return tfs
}

// The function that actually counts the most.
func DriftReport(absProjectPath string, tfBinary string) *TerraformService {
	// terraform init
	service := ConfigureTerraform(absProjectPath, tfBinary)
	Init(service)

	// terraform show
	state := Show(service)

	// terraform plan (-detailed-exitcode)
	var isPlanned bool = Plan(service)

	// terraform plan (-out=out.tfplan)
	planPath := fmt.Sprintf("%s/out.tfplan", absProjectPath)
	var rawPlan = ShowPlanFileRaw(service, planPath)
	planString := GetResourceModificationCount(rawPlan)

	// If drift detected in Plan return the Add/Change/Destroy count values.
	modifiedResourceCount := ParseResourceModificationCount(planString)

	// Get project name + status information
	_, projectName := GetProjectName(absProjectPath)
	summary := GetDriftSummary(isPlanned, state)

	// Format a TerraformService structure with all information needed for the Drift Report
	tfService := UpdateDriftReportData(state, projectName, modifiedResourceCount, summary)
	return tfService
}

// Go channel which returns the result of a DriftReport (required to parallelize)
func GetProjectDrift(ch chan *TerraformService, absProjectPath string, tfBinary string) {
	log.Printf("[GetDriftReport] Getting values for project: %s", absProjectPath)
	ch <- DriftReport(absProjectPath, tfBinary)
}
