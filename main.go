package main

import (
	"Fisherman/agent/info"
	"Fisherman/server"
	"fmt"
)

func main() {
	hostname := info.GetHostname()
	fmt.Printf("Entering bad waters, Fisherman is lurking in %s.\n", hostname)

	server.Setup()
}
