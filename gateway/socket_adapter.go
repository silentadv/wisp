package gateway

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

type SocketAdapter struct {
	Conn *websocket.Conn
}

func NewSocketAdapter(url string) (*SocketAdapter, error) {
	header := http.Header{}
	conn, _, err := websocket.DefaultDialer.Dial(url, header)
	if err != nil {
		return nil, err
	}

	adapter := &SocketAdapter{
		Conn: conn,
	}

	return adapter, nil
}

func (s *SocketAdapter) Listen(callback func(*GatewayReceivePayload)) {
	defer s.Conn.Close()

	for {
		_, message, err := s.Conn.ReadMessage()
		if err != nil {
			fmt.Printf("wisp (error) -> socket-adapter failed to read wss message: %s.\n", err.Error())
			return
		}

		payload, err := ParseIncomingPayload(message)
		if err != nil {
			fmt.Printf("wisp (error) -> socket-adapter failed to parse wss payload: %s.\n", err.Error())
			return
		}

		callback(payload)

		fmt.Printf("wisp -> socket-adapter received message: %s\n", string(message))
	}
}

func (s *SocketAdapter) Send(payload GatewaySendPayload) error {
	err := s.Conn.WriteJSON(payload)
	if err != nil {
		return fmt.Errorf("wisp (error) -> socket-adapter failed to send message: %s\n", err.Error())
	}

	return nil
}

