package terraform

import (
	"path/filepath"

	tfjson "github.com/hashicorp/terraform-json"

	"clairvoyance/app/reporting"
	"clairvoyance/log"
)

// Default TerraformService struct for clairvoyance reporting.
type TerraformService struct {
	//State            *tfjson.State `json:"state"`
	ProjectName      string `json:"project_name"`
	TerraformVersion string `json:"terraform_version"`
	CountAdd         int    `json:"count_add"`
	CountChange      int    `json:"count_change"`
	CountDestroy     int    `json:"count_destroy"`
	Summary          string `json:"summary"`
}

// Retrieve full file path to the project's terraform.tfstate
func GetStateFile(tfProjectPath string) string {
	var statefile string = tfProjectPath + "/.terraform/terraform.tfstate"
	log.Printf("[GetStateFile] %s", statefile)
	return statefile
}

//Split absolute path into root + directory
func GetProjectName(projectName string) (string, string) {
	paths := []string{projectName}

	for _, p := range paths {
		dir, file := filepath.Split(p)
		return dir, file
	}
	return "", ""
}

// For each Terraform resource print the address
//var resources []*tfjson.StateResource = state.Values.RootModule.Resources
func ResourceAddressList(state *tfjson.State) {
	var resources []*tfjson.StateResource = state.Values.RootModule.Resources
	var resourceMap map[string][]byte
	resourceMap = make(map[string][]byte)

	for i, res := range resources {
		resourceValues := reporting.FormatTerraformResource(resources[i])
		resourceMap[res.Address] = resourceValues
	}
}
