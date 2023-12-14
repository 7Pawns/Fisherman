package main

import (
	"Fisherman/agent/info"
	"Fisherman/server"
	"Fisherman/server/baits"
	"fmt"
	"os/exec"
)

func main() {
	hostname := info.GetHostname()
	fmt.Printf("Entering bad waters, Fisherman is lurking in %s.\n", hostname)

	bait := baits.NewBait("C:\\Users\\orads\\Downloads\\simple_http.ps1")

	exec.Command("Powershell", "Invoke-Expression", bait.BaitToString("localhost:8080/hook"))
	server.Setup()
}
