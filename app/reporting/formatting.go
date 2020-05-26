package reporting

import (
	"fmt"

	"clairvoyance/log"
)

func FormatCodeBlock(message string) string {
	return fmt.Sprintf("\n```\n%s\n```", message)
}

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

func DebugFormatMessage() string {
	formattedMessage := "Hello world!\n" +
		"```" +
		"code\n" +
		"block\n" +
		"```" +
		"\n\n" +
		"> who me?" +
		"\n\n" +
		"yes, *YOU*!"

	log.Printf("formatting - Debug format message:\n%s", formattedMessage)
	return formattedMessage
}
