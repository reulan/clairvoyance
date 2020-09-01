package terraform

import (
	"context"

	tfexec "github.com/hashicorp/terraform-exec/tfexec"
	tfinstall "github.com/hashicorp/terraform-exec/tfinstall"
	tfjson "github.com/hashicorp/terraform-json"

	"clairvoyance/log"
)

var TerraformContext = context.Background()

func ConfigureTerraform(workingDir string) *tfexec.Terraform {
	// Generate a new tfinstall Terraform configuration
	execPath, err := tfinstall.Find()
	if err != nil {
		log.Errorf("cli/ConfigureTerraform - %s", err)
	}

	service, err := tfexec.NewTerraform(workingDir, execPath)
	if err != nil {
		log.Errorf("cli/ConfigureTerraform - %s", err)
	}

	log.Info("cli/ConfigureTerraform - Created tfexec configuration.")

	return service
}

func Init(service *tfexec.Terraform) *tfexec.Terraform {
	// Run `terraform init` so that the working directories context can be initialized.
	log.Info("cli/Init - terraform init")

	err := service.Init(context.Background(), tfexec.Upgrade(true), tfexec.LockTimeout("60s"))
	if err != nil {
		log.Errorf("cli/Init - %s", err)
	}

	return service
}

func Show(service *tfexec.Terraform) *tfjson.State {
	// Run `terraform show` against the state defined in the working directory.
	log.Info("cli/Show - terraform show")

	state, err := service.Show(context.Background())
	if err != nil {
		log.Errorf("cli/Show - %s", err)
	}

	return state
}
