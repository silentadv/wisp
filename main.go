package main

import (
	"fmt"
	"os"

	"github.com/silentadv/wisp/gateway"
	"github.com/silentadv/wisp/internals"
)

func main() {
	shard := gateway.NewShard(0, os.Getenv("WISP_DISCORD_TOKEN"), 0, 1)
	shard.Connect()

	if err := internals.StartDaemon("localhost:4000"); err != nil {
		fmt.Printf("wisp (error) -> failed to start daemon: %v\n", err)
		return
	}
}

