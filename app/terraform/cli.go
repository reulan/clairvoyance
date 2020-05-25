package terraform

import (
	"fmt"
	"os"

	"github.com/kmoe/terraform-exec/tfexec"
	tfjson "github.com/hashicorp/terraform-json"

	"clairvoyance/log"
)


func Configure(ExecPath string, WorkingDir string) tfexec.Terraform{} {
	// Override parameters for testing
	ExecPath := ""
	WorkingDir := os.Getenv("GOPATH") + "/src/github.com/kmoe/terraform-exec/testdata"

	// Generate a new tfexec Terraform configuration
	tfcfg, err := tfexec.NewTerraform(ExecPath, WorkingDir)

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

func Plan(tfcfg tfexec.Terraform) *tfjson.State {
	// Run `terraform plan` against the state defined in the working directory.
	log.Info("Clarivoyance - terraform plan on {%s}", tfcfg.)
	state, err := cfg.Plan()
	if err != nil {
		panic(err)
	}

	// Print all returned values from the `terraform plan` command (of type *tfjson.State)
	fmt.Println(state.FormatVersion) // "0.1"
	fmt.Println(state.TerraformVersion)
	fmt.Println(state.Values)

	return state
}
