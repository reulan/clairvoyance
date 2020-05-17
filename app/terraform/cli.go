package terraform

import (
	"fmt"
	"os"

	tfexec "github.com/kmoe/terraform-exec"
	tfjson "github.com/hashicorp/terraform-json"
)

func Configure() tfexec.Config {
	workingDir := os.Getenv("GOPATH") + "/src/github.com/kmoe/terraform-exec/testdata"
	cfg := tfexec.Config{
		WorkingDir: workingDir,
	}
	return cfg
}

func Init(cfg tfexec.Config) {
	// Run `terraform init` so that the working directories state can be initialized.
	fmt.Println("Clarivoyance - terraform init")
	err := cfg.Init()
	if err != nil {
		panic(err)
	}
}

func Show(cfg tfexec.Config) *tfjson.State {
	// Run `terraform show` against the state defined in the working directory.
	fmt.Println("Clarivoyance - terraform show")
	state, err := cfg.Show()
	if err != nil {
		panic(err)
	}

	// Print all returned values from the `terraform show` command (of type *tfjson.State)
	fmt.Println(state.FormatVersion) // "0.1"
	fmt.Println(state.TerraformVersion)
	fmt.Println(state.Values)

	return state
}
