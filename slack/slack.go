package slack

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
)

const (
	apiURL  = "https://hooks.slack.com/services"
	webhook = "/T207WP78B/B012XGPKUSZ/cKhZjr5JhWRmkp2bGX9qPQQ0"
)

// Slack .
type Slack struct{}

// Notify .
func (s *Slack) Notify(data string) {
	sendMessage(data)
}

func sendMessage(data string) {
	msg := fmt.Sprintf(`{"text":%q}`, data)
	req, err := http.NewRequest(http.MethodPost, apiURL+webhook, bytes.NewReader([]byte(msg)))

	if err != nil {
		log.Fatalf("It couldnt create a new request : %v", err)
	}

	client := http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Fatalf("It couldn't send a new message via slack: %v", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Fatalf("The error status is an unexpected error: %d", resp.StatusCode)
	}

	fmt.Println("Message sent via slack")
}
