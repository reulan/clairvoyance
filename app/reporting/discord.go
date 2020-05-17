package reporting

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"

	"clairvoyance/log"
)

var (
    // Sends a message to #clairvoyance in the noobshack discord.
    DiscordWebhookName string = "#clairvoyance in noobshack"
    DiscordWebhookSecret string = os.Getenv("DISCORD_WEBHOOK_SECRET")
	DiscordWebhookURL string = "https://discordapp.com/api/webhooks/" + DiscordWebhookSecret
	ContentType string = "application/json"
)

// TODO: Gather then format Terraform drift report results
//var report = map[string]string{"content": "hello world from Clairvoyance! content!", "message": "hw from cv - message!"}
//var reportJSON, _ := json.Marshal(report)
//var reportString string = "Clairvoyance says hello to this world."

func SendReport() {
	// Populate the JSON payload and Marshall data for request
	reportString := map[string]string{
		"content": "Hello discord!",
		"username": "clairvoyance",
	}
	reportJSON, err := json.Marshal(reportString)

	// Create the HTTP socket
	timeout := time.Duration(10 * time.Second)
	client := &http.Client{
		Timeout: timeout,
	}

	// Configure payload parameters
	params := url.Values{}
	params.Set("wait", "true")

	// Format and send the HTTP request
	request, err := http.NewRequest("POST", DiscordWebhookURL, bytes.NewBuffer(reportJSON))
	request.Header.Set("Content-Type", ContentType)
	request.Header.Set("X-Custom-Header", "clairvoyance")
	request.Header.Set("Content-Length", strconv.Itoa(len(params.Encode())))
	request.URL.RawQuery = params.Encode()

	// Send message to Discord
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}

	// Close the request before reporting or returning data
	defer response.Body.Close()

	// Log metrics and information
	log.Printf("Sent message to: %s - Status: [%s]\n", DiscordWebhookName, response.Status)

	// Extract body of response and read the contents.
	body, _ := ioutil.ReadAll(response.Body)
	log.Debug("Body: ", string(body))
}
