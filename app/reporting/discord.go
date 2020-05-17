package reporting

import (
	"net/http"
)

var (
    // Sends a message to #clairvoyance in the noobshack discord.
    discordWebhookURL string = "https://discordapp.com/api/webhooks/711450720179585025/JCVcCoi_trb6JEEGKqNWuPdpHXQRbcPQsK3TFKW7yb5JusZSPg1exw8gfZQQkYUEWBeB"
	contentType string = "application/json"
)

func SendMessage() {
    resp, err := http.Post(discordWebhookURL, contentType)
}
