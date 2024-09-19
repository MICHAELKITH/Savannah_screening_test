package models
// Recipient is a model
type Recipient struct {
	Number    string `json:"number"`
	Cost      string `json:"cost"`
	Status    string `json:"status"`
	StatusCode  int64 `json:"statusCode"`
	MessageID string `json:"messageId"`
}

type SendMessageResponse struct {
	SMS SMS2 `json:"SMSMessageData"`
}

// SMS2 is a model
type SMS2 struct {
	Message string `json:"Message"`
	Recipients []Recipient `json:"recipients"`
}