package request

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var webhooks *WebhookInfo

func Qase(channel string, payload interface{}) {
	Read()
	for _, v := range webhooks.Qase {
		if v.Channel == channel {
			requestBody, err := json.Marshal(payload)
			if err != nil {
				fmt.Printf("Qase Marshal error: %v\n", err)
				return
			}
			res, err := http.Post(v.URL, "application/json", bytes.NewBuffer(requestBody))
			if err != nil {
				fmt.Printf("Qase request error: %v\n", err)
				return
			}
			defer res.Body.Close()
			break
		}
	}
}

func Read() {
	if webhooks == nil {
		rawJson, err := os.Open("./webhooks.json")
		if err != nil {
			log.Fatalf("Failed to read webhooks.json file: %v", err)
		}
		defer rawJson.Close()

		rawByteJson, _ := ioutil.ReadAll(rawJson)
		json.Unmarshal(rawByteJson, &webhooks)
	}
}
