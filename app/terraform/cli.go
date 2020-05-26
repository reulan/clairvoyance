package terraform

import (
	"fmt"
	"os"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/kmoe/terraform-exec/tfexec"

	"clairvoyance/log"
)

var WorkingDir = os.Getenv("GOPATH") + "/src/github.com/kmoe/terraform-exec/testdata"

// Override parameters for testing
// WorkingDir string

func ConfigureTerraform() *tfexec.Terraform {
	// Generate a new tfexec Terraform configuration
	execPath, err := tfexec.FindTerraform()
	if err != nil {
		panic(err)
	}

	tfcfg, err := tfexec.NewTerraform(execPath, WorkingDir)
	if err != nil {
		panic(err)
	}

	log.Info("Clarivoyance - Created tfexec configuration.")
	return tfcfg
}

func Init(tfcfg tfexec.Terraform) {
	// Run `terraform init` so that the working directories state can be initialized.
	log.Info("Clarivoyance - terraform init")
	err := tfcfg.Init()
	if err != nil {
		panic(err)
	}
}

func Show(tfcfg tfexec.Terraform) *tfjson.State {
	// Run `terraform show` against the state defined in the working directory.
	log.Info("Clarivoyance - terraform show")
	state, err := tfcfg.Show()
	if err != nil {
		panic(err)
	}

	// Print all returned values from the `terraform show` command (of type *tfjson.State)
	fmt.Println(state.FormatVersion)
	fmt.Println(state.TerraformVersion)
	fmt.Println(state.Values)

	return state
}

func Plan(tfcfg tfexec.Terraform) {
	// Run `terraform plan` against the state defined in the working directory.
	log.Info("Clarivoyance - terraform plan on {%s}", tfcfg.)

	// Currently unimplemented in tfexec
	err := tfcfg.Plan()
	if err != nil {
		panic(err)
	}
}

func Apply(tfcfg tfexec.Terraform) {
	// Run `terraform apply` against the state defined in the working directory.
	log.Info("Clarivoyance - terraform apply on {%s}", tfcfg.)

	// Currently unimplemented in tfexec
	err := tfcfg.Apply()
	if err != nil {
		panic(err)
	}
}
