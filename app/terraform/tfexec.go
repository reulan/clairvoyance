package terraform

import (
	"context"

	tfexec "github.com/hashicorp/terraform-exec/tfexec"
	tfjson "github.com/hashicorp/terraform-json"

	"clairvoyance/log"
)

var TerraformContext = context.Background()

func ConfigureTerraform(workingDir string, execPath string) *tfexec.Terraform {
	execPath = "/usr/bin/terraform"
	service, err := tfexec.NewTerraform(workingDir, execPath)
	if err != nil {
		panic(err)
	}
	log.Printf("[cli/tfexec/ConfigureTerraform] Created tfexec configuration for project: %s.", workingDir)
	return service
}

// Run `terraform init` so that the working directories context can be initialized.
func Init(service *tfexec.Terraform) {
	err := service.Init(TerraformContext)
	if err != nil {
		panic(err)
	}
	log.Info("[cli/tfexec/Init] Initialized Terraform project.")
}

// Run `terraform show` against the state defined in the working directory.
func Show(service *tfexec.Terraform) *tfjson.State {
	state, err := service.Show(TerraformContext)
	if err != nil {
		panic(err)
	}
	return state
}

// Run `terraform plan` against the state defined in the working directory.
// 0 = false (no changes)
// 2 = true  (drift)
func Plan(service *tfexec.Terraform) bool {
	log.Info("[cli/tfexec/Plan] - terraform plan")
	isPlanned, err := service.Plan(TerraformContext)
	//isPlanned, err := service.Plan(TerraformContext, opts ...PlanOptions)
	if err != nil {
		panic(err)
	}
	return isPlanned
}

// Run `terraform plan` against the state defined in the working directory.
// 0 = false (no changes)
// 2 = true  (drift)
func PlanWithOptions(service *tfexec.Terraform, outPath string) bool {
	log.Info("[cli/tfexec/PlanWithOptions] - terraform plan")
	isPlanned, err := service.Plan(TerraformContext)
	//isPlanned, err := service.Plan(TerraformContext, opts ...PlanOptions)
	if err != nil {
		panic(err)
	}
	return isPlanned
}

// Run `terraform plan` against the state defined in the working directory.
func ShowPlanFileRaw(service *tfexec.Terraform, planPath string) string {
	log.Info("[cli/tfexec/PlanFileRaw] terraform plan (human readable)")
	plan, err := service.ShowPlanFileRaw(TerraformContext, planPath)
	if err != nil {
		panic(err)
	}
	return plan
}
