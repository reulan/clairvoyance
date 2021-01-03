package terraform

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"

	"clairvoyance/app/reporting"

	tfjson "github.com/hashicorp/terraform-json"
	tail "github.com/hpcloud/tail"
)

// Default TerraformService struct for clairvoyance reporting.
type TerraformService struct {
	//State            *tfjson.State `json:"state"`
	ProjectName      string `json:"project_name"`
	TerraformVersion string `json:"terraform_version"`
	CountAdd         int    `json:"count_add"`
	CountChange      int    `json:"count_change"`
	CountDestroy     int    `json:"count_destroy"`
	Summary          string `json:"summary"`
}

// Retrieve full file path to the project's terraform.tfstate
func GetStateFile(tfProjectPath string) string {
	var statefile string = tfProjectPath + "/.terraform/terraform.tfstate"
	log.Printf("[GetStateFile] %s", statefile)
	return statefile
}

//func ExtractDriftReportData(state *tfjson.State, projectName string, counts map[string]int, summary string) *TerraformService {
func ExtractDriftReportData(state *tfjson.State, projectName string, counts map[string]int, summary string) string {
	tfs := &TerraformService{
		//State:            state,
		ProjectName:      projectName,
		TerraformVersion: state.TerraformVersion,
		CountAdd:         counts["CountAdd"],
		CountChange:      counts["CountChange"],
		CountDestroy:     counts["CountDestroy"],
		Summary:          summary,
	}

	tfsb, _ := json.Marshal(tfs)
	var message string = fmt.Sprintf("```%s```", tfsb)
	log.Printf("[ExtractDriftReportData] TerraformService struct is:\n%+v", tfs)

	return message
}

// true == diff || false == No changes.
func DriftDetection(exitStatus bool, state *tfjson.State) string {
	var message string
	if exitStatus {
		message = "Drift detected for Plan."
		log.Printf("[DriftDetection] %s", message)
	} else if !exitStatus {
		message = "No changes."
		log.Printf("[DriftDetection] %s", message)
	} else {
		message = "Error planning Terraform project."
		log.Printf("[DriftDetection] %s", message)
	}
	return message
}

// For each Terraform resource print the address
//var resources []*tfjson.StateResource = state.Values.RootModule.Resources
func ResourceAddressList(state *tfjson.State) {
	var resources []*tfjson.StateResource = state.Values.RootModule.Resources
	var resourceMap map[string][]byte
	resourceMap = make(map[string][]byte)

	for i, res := range resources {
		resourceValues := reporting.FormatTerraformResource(resources[i])
		resourceMap[res.Address] = resourceValues
	}
}

func ResourceModificationCount(planFileRawString string) string {
	var filename string = "/tmp/clairvoyance-tmp"
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	} else {
		file.WriteString(planFileRawString)
		log.Printf("Wrote %s as staging for tfplan conversion.", filename)
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

	//TODO: fix me potential infinite loop
	return ""
}

func ParseModificationCount(resourceModificationString string) map[string]int {
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
