package gateway

type Shard struct {
	Id         int
	Token      string
	Intents    int
	ShardCount int
	socket     *SocketAdapter
}

const (
	gatewayURL = "wss://gateway.discord.gg/?v=10&encoding=json"
)

func NewShard(id int, token string, intents int, shardCount int) *Shard {
	return &Shard{
		Id:         id,
		Token:      token,
		Intents:    intents,
		ShardCount: shardCount,
	}
}

func (s *Shard) Connect() {
	socket, err := NewSocketAdapter(gatewayURL)
	if err != nil {
		return
	}

	s.socket = socket
	s.socket.Listen(s.handleIncomingMessage)
}

func (s *Shard) handleIncomingMessage(p *GatewayReceivePayload) {
	if p.Op == 10 {
		s.identify()
		return
	}
}

func (s *Shard) identify() {
	properties := GatewayIdentifyProperties{
		Os:      "linux",
		Browser: "wisp",
		Device:  "wisp",
	}

	identify := GatewayIdentifyData{
		Intents:    s.Intents,
		Shard:      [2]int{s.Id, s.ShardCount},
		Token:      s.Token,
		Properties: properties,
	}

	payload := GatewaySendPayload{
		Op: 2,
		D:  identify,
	}

	s.socket.Send(payload)
}
