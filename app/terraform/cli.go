package terraform

import (
	"context"
	"fmt"

	tfexec "github.com/hashicorp/terraform-exec/tfexec"
	tfinstall "github.com/hashicorp/terraform-exec/tfinstall"
	tfjson "github.com/hashicorp/terraform-json"

	"clairvoyance/log"
)

var TerraformContext = context.Background()

func main() {
	binaryPath = 
	workingDir := "/path/to/working/dir"
	tf, err := tfexec.NewTerraform(workingDir, execPath)
	if err != nil {
		panic(err)
	}

	err = tf.Init(context.Background(), tfexec.Upgrade(true), tfexec.LockTimeout("60s"))
	if err != nil {
		panic(err)
	}

	state, err := tf.Show(context.Background())
	if err != nil {
		panic(err)
	}

	fmt.Println(state.FormatVersion) // "0.1"
}

func ConfigureTerraform(workingDir string) *tfexec.Terraform {
	// Generate a new tfexec Terraform configuration
	execPath, err := tfexec.FindTerraform()
	if err != nil {
		log.Errorf("cli/ConfigureTerraform - %s", err)
	}

	service, err := tfexec.NewTerraform(execPath, workingDir)
	if err != nil {
		log.Errorf("cli/ConfigureTerraform - %s", err)
	}

	log.Info("cli/ConfigureTerraform - Created tfexec configuration.")

	return service
}

func Init(service *tfexec.Terraform) {
	// Run `terraform init` so that the working directories context can be initialized.
	log.Info("cli/Init - terraform init")

	err := service.Init(TerraformContext)
	if err != nil {
		log.Errorf("cli/Init - %s", err)
	}
}

func Show(service *tfexec.Terraform) *tfjson.State {
	// Run `terraform show` against the state defined in the working directory.
	log.Info("cli/Show - terraform show")

	state, err := service.Show(TerraformContext)
	if err != nil {
		log.Errorf("cli/Show - %s", err)
	}

	return state
}
