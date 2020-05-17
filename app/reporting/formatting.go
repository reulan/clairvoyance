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
	formattedMessage := "Syntax check:" +
		"\n" +
		"```" +
		"code" +
		"block" +
		"```" +
		"> indent" +
		"*bold*" +
		"_italics_" +
		"\n - do you need to new line?" +
		"`single line code` " +
		"```" +
		"I used to be plat" +
		"\n" +
		"but now I am gold" +
		"```"

	log.Printf("formatting - Debug format message:\n%s", formattedMessage)
	return formattedMessage
}

