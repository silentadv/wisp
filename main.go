package main 

import (
	"github.com/silentadv/wisp/internals"
	"fmt"
)

func main() {
	if err := internals.StartDaemon("localhost:4000"); err != nil {
		fmt.Printf("wisp (error) -> failed to start daemon: %v\n", err)
		return
	}
}