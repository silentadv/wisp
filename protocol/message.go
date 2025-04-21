package protocol

import (
	"encoding/json"
)

type Message struct {
	Op   string `json:"op"`
	Data string `json:"d"`
}

func ParseMessage(raw []byte) (*Message, error) {
	var m Message
	err := json.Unmarshal(raw, &m)
	return &m, err
}

