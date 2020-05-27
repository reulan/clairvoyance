package reporting

import (
	"clairvoyance/log"
)

func SendMessageStdout(message string) {
	// Log metrics and information
	log.Printf("Formatted message:\n%s", message)
}

func SendJSONStdout(message []byte) {
	// Log metrics and information
	log.Printf("Formatted message:\n%s", message)
}
