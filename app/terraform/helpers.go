package terraform

import (
	"log"
	"strings"

	tfjson "github.com/hashicorp/terraform-json"
)

// Retrieve full file path to the project's terraform.tfstate
func GetStateFile(tfProjectPath string) string {
	var statefile string = tfProjectPath + "/.terraform/terraform.tfstate"
	log.Printf("[GetStateFile] %s", statefile)
	return statefile
}

// true == diff || false == No changes.

func DriftDetection(exitStatus bool, state *tfjson.State) {
	var messages []string
	if exitStatus {
		log.Printf("[DriftDetection] Drift detected for plan.")
		var resources []*tfjson.StateResource = state.Values.RootModule.Resources
		messages = ResourceAddressList(resources)
		message := strings.Join(messages, "\n")
		log.Println(message)

	} else if !exitStatus {
		log.Printf("[DriftDetection] No changes.")
	} else {
		log.Printf("[DriftDetection] Error planning Terraform project.")
	}
}

// For each Terraform resource print the address
//var resources []*tfjson.StateResource = state.Values.RootModule.Resources
func ResourceAddressList(resources []*tfjson.StateResource) []string {
	var messages []string
	for _, res := range resources {
		messages = append(messages, res.Address)
		log.Printf("added %s to messages", res.Address)
	}
	return messages
}
