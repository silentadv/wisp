package gateway

import "encoding/json"

type GatewayReceivePayload struct {
	Op int             `json:"op"`
	D  json.RawMessage `json:"d"`
	S  *int            `json:"s,omitempty"`
	T  *string         `json:"t,omitempty"`
}

type GatewaySendPayload struct {
	Op int `json:"op"`
	D  any `json:"d"`
}

type GatewayIdentifyData struct {
	Intents    int                       `json:"intents"`
	Shard      [2]int                    `json:"shard"`
	Token      string                    `json:"token"`
	Properties GatewayIdentifyProperties `json:"properties"`
}

type GatewayIdentifyProperties struct {
	Os      string `json:"os"`
	Browser string `json:"browser"`
	Device  string `json:"device"`
}

func ParseIncomingPayload(raw []byte) (*GatewayReceivePayload, error) {
	var p GatewayReceivePayload
	err := json.Unmarshal(raw, &p)
	return &p, err
}
