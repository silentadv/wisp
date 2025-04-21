package internals

import (
	"fmt"
	"github.com/silentadv/wisp/protocol"
	"net"
)

func StartDaemon(address string) error {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return fmt.Errorf("wisp (error) -> failed to start daemon: %w\n", err)
	}

	fmt.Println("wisp -> daemon started, waiting for connections...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("wisp (error) -> failed to accept connection: %v\n", err)
			continue
		}

		fmt.Println("wisp -> new connection accepted")
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Printf("wisp (error) -> failed to read from connection: %v\n", err)
		return
	}

	msg, err := protocol.ParseMessage(buf[:n])
	if err != nil {
		fmt.Printf("wisp (error) -> failed to parse message: %v\n", err)
		return
	}

	fmt.Printf("wisp -> received message: %s\n", msg)
}

