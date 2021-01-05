package terraform

import (
	"fmt"
	"log"
)

///tfChan, absProjectPath, tfBinary
func GetProjectDrift(ch chan *TerraformService, absProjectPath string, tfBinary string) {
	log.Println("[GetDriftReport] Getting values for project...")
	ch <- DriftReport(absProjectPath, tfBinary)
}

func DriftReport(absProjectPath string, tfBinary string) *TerraformService {
	// terraform init
	service := ConfigureTerraform(absProjectPath, tfBinary)
	Init(service)

	// terraform show
	state := Show(service)

	// terraform plan
	// TODO: tf plan (with -out=out.tfplan)
	planOptions := fmt.Sprintf("-out=%s/out.tfplan", absProjectPath)
	var po []string = []string{planOptions}
	fmt.Println(po)
	var isPlanned bool = Plan(service)
	planPath := fmt.Sprintf("%s/out.tfplan", absProjectPath)
	var rawPlan = ShowPlanFileRaw(service, planPath)
	//log.Printf("rawPlan: %s", rawPlan)
	planString := ResourceModificationCount(rawPlan)
	modifiedResourceCount := ParseModificationCount(planString)
	summary := DriftDetection(isPlanned, state)

	_, projectName := GetProjectName(absProjectPath)
	tfService := ExtractDriftReportData(state, projectName, modifiedResourceCount, summary)
	return tfService
}
