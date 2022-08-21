package models

type WebsocketMesg struct {
	Type    string      `json:"type"`
	Payload interface{} `json:"payload"`
}
