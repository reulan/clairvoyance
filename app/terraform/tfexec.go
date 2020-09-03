package terraform

import (
	"context"

	tfexec "github.com/hashicorp/terraform-exec/tfexec"
	tfjson "github.com/hashicorp/terraform-json"

	"clairvoyance/log"
)

var TerraformContext = context.Background()

func ConfigureTerraform(workingDir string, execPath string) *tfexec.Terraform {
	service, err := tfexec.NewTerraform(workingDir, execPath)
	if err != nil {
		log.Errorf("cli/ConfigureTerraform - %s", err)
	}

	log.Info("cli/ConfigureTerraform - Created tfexec configuration.")

	return service
}

// Run `terraform init` so that the working directories context can be initialized.
func Init(service *tfexec.Terraform) {
	log.Info("cli/Init - terraform init")

	err := service.Init(TerraformContext)
	if err != nil {
		log.Errorf("cli/Init - %s", err)
	}
}

// Run `terraform show` against the state defined in the working directory.
func Show(service *tfexec.Terraform) *tfjson.State {
	log.Info("cli/Show - terraform show")

	state, err := service.Show(TerraformContext)
	if err != nil {
		log.Errorf("cli/Show - %s", err)
	}

	return state
}

// Run `terraform plan` against the state defined in the working directory.
// 0 = false (no changes)
// 2 = true  (drift)
func Plan(service *tfexec.Terraform) bool {
	log.Info("cli/Plan - terraform plan")

	hasChanges, err := service.Plan(TerraformContext)
	if err != nil {
		log.Errorf("cli/Plan - %s", err)
	}

	return hasChanges
}
