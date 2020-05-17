package reporting

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"clairvoyance/log"
)

var (
    // Sends a message to #clairvoyance in the noobshack discord.
    discordWebhookName string = "#clairvoyance in noobshack"
    discordWebhookSecret string = os.Getenv("DISCORD_WEBHOOK_SECRET")
	discordWebhookURL string = "https://discordapp.com/api/webhooks/" + discordWebhookSecret
	contentType string = "application/json"
)

// Gather report results and return them as a byte array in order to be shipped off
var report = []byte(`{"report":"hello world from clairvoyance!"}`)

func SendReport() {
	request, err := http.NewRequest("POST", discordWebhookURL, bytes.NewBuffer(report))
	request.Header.Set("X-Custom-Header", "clairvoyance")
	request.Header.Set("Content-Type", contentType)

	// Create the HTTP socket and send request
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}

	// Close the request before reporting or returning data
	defer response.Body.Close()

	// Log metrics and information
	log.Info(fmt.Printf("Sent message to: %s", discordWebhookName))
	log.Info(fmt.Printf("HTTP Status code: %s", response.Status ))
	log.Info(fmt.Printf("Headers: %s", response.Header))

	// Extract body of response and read the contents.
	body, _ := ioutil.ReadAll(response.Body)
	log.Info(fmt.Println("Body:", string(body)))
}
