package request

type WebhookInfo struct {
	Qase []struct {
		Channel string `json:"channel"`
		URL     string `json:"url"`
	} `json:"qase"`
}
