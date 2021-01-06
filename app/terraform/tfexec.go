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
	log.Debugf("[ConfigureTerraform] Created tfexec configuration for project: %s.", workingDir)
	return service
}

// Run `terraform init` so that the working directories context can be initialized.
func Init(service *tfexec.Terraform) {
	err := service.Init(TerraformContext, tfexec.Lock(false))
	if err != nil {
		panic(err)
	}
	log.Debug("[Init] Initialized Terraform project.")
}

// (-detailed-exitcode)
// Run `terraform plan` against the state defined in the working directory.
// 0 = false (no changes)
// 2 = true  (drift)
func Plan(service *tfexec.Terraform) bool {
	isPlanned, err := service.Plan(TerraformContext, tfexec.Out("out.tfplan"), tfexec.Lock(false))
	if err != nil {
		panic(err)
	}
	log.Debug("[Plan] Planning Terraform service and writing to out.tfplan.")
	return isPlanned
}

// View State after it's been initialized and refreshed
// Run `terraform show` against the state defined in the working directory.
func Show(service *tfexec.Terraform) *tfjson.State {
	state, err := service.Show(TerraformContext)
	if err != nil {
		panic(err)
	}
	log.Debug("[Show] Retrieving state object for project.")
	return state
}

// Run `terraform plan` against the state defined in the working directory.
func ShowPlanFileRaw(service *tfexec.Terraform, planPath string) string {
	log.Debug("[ShowPlanFileRaw] Human readable Plan derived from out.tfplan.")
	plan, err := service.ShowPlanFileRaw(TerraformContext, planPath)
	if err != nil {
		panic(err)
	}
	return plan
}
