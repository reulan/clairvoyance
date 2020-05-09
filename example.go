package main

import (
	"fmt"
	"os"

	tfexec "github.com/kmoe/terraform-exec"
)

func main() {
	workingDir := os.Getenv("GOPATH") + "/src/github.com/kmoe/terraform-exec/testdata"

	// Run `terraform init` so that the working directories state can be initialized.
	err := tfexec.Init(workingDir)
	if err != nil {
		panic(err)
	}

	// Run `terraform show` against the state defined in the working directory.
	state, err := tfexec.Show(workingDir)
	if err != nil {
		panic(err)
	}

	// Print all returned values from the `terraform show` command (of type *tfjson.State)
	fmt.Println(state.FormatVersion) // "0.1"
	fmt.Println(state.TerraformVersion)
	fmt.Println(state.Values)
}
