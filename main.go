package main

import (
	"Fisherman/agent/info"
	"Fisherman/server"
	"Fisherman/server/baits"
	"fmt"
)

func main() {
	hostname := info.GetHostname()
	fmt.Printf("Entering bad waters, Fisherman is lurking in %s.\n", hostname)

	bait := baits.NewBait("C:\\path\\to\\simple_http.ps1")

	fmt.Print(bait.GetBaitScript("localhost:8080/hook"))

	server.Setup()
}
