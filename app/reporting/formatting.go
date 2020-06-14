package reporting

import (
	"encoding/json"
	"fmt"

	tfjson "github.com/hashicorp/terraform-json"

	"clairvoyance/log"
)

func FormatDriftReport(message string) string {
	var formattedMessage string = fmt.Sprintf("Terraform Drift Detection Report"+
		"\n"+
		"```"+
		"%s"+
		"```",
		message)

	log.Printf("formatting - Formatted message.\n%s", formattedMessage)
	return formattedMessage
}

// Terraform Output
func FormatTerraformShow(state *tfjson.State) []byte{
	output, err := json.MarshalIndent(state, "", "\t")
	if err != nil {
		log.Errorf("reporting/formatting - %s", err)
	}

	return output
}
