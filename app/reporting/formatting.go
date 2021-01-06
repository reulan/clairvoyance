package reporting

import (
	"encoding/json"
	"fmt"

	tfjson "github.com/hashicorp/terraform-json"

	"clairvoyance/log"
)

// eventually have stdout template and discord template
func FormatDriftReport(message string) string {
	var title string = "Terraform Drift Detection Report"
	var formattedMessage string = fmt.Sprintf("%s\n```\n%s\n```", title, message)
	log.Debugf("[FormatDriftReport] Formatted message.\n%s", formattedMessage)
	return formattedMessage
}

// Terraform Output
func FormatTerraformShow(state *tfjson.State) []byte {
	output, err := json.MarshalIndent(state, "", "\t")
	if err != nil {
		log.Errorf("[FormatTerraformShow] %s", err)
	}
	log.Debugf("[FormatTerraformShow] Formatted message:\n%s", output)
	return output
}

func FormatTerraformResource(resource *tfjson.StateResource) []byte {
	formattedResource, err := json.MarshalIndent(resource, "", "\t")
	if err != nil {
		log.Errorf("[FormatTerraformResources] %s", err)
	}
	log.Debugf("[FormatTerraformResource] Formatted message:\n%s", formattedResource)
	return formattedResource
}
