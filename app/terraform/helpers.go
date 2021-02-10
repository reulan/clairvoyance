package terraform

import (
	"os"
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

// Clean up cached Terraform project files
func CleanupCachedFiles(serviceDir string) {
	//serviceDir, _ = filepath.Abs(serviceDir)
	// If serviceDir != absolute path, use CVPD + serviceDir
	//os.Getenv("CLAIRVOYANCE_PROJECT_DIR")

	// rm -rf ./.terraform
	var terraformInitDir string = (serviceDir + "/.terraform")
	log.Debugf("[CleanUpCachedFiles] DELETING: %s", terraformInitDir)
	os.RemoveAll(terraformInitDir)

	// rm terraform.lock.hcl
	var terraformLockFile string = (serviceDir + "/.terraform.lock.hcl")
	log.Debugf("[CleanUpCachedFiles] DELETING: %s", terraformLockFile)
	os.Remove(terraformLockFile)
}

func DetectTerraformVersion() string {
	// Setup Terraform Version to use
	var terraformVersion string
	var _, tfVersionSet = os.LookupEnv("CLAIRVOYANCE_TERRAFORM_VERSION")

	if tfVersionSet {
		terraformVersion = os.Getenv("CLAIRVOYANCE_TERRAFORM_VERSION")
	}

	terraformVersion = "0.14.5"
	return terraformVersion
}
