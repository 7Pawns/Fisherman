package main

import (
	"ShyGuy/agent/info"
	"ShyGuy/server"
	"fmt"
)

func main() {
	hostname := info.GetHostname()
	fmt.Printf("ShyGuy initialized on %s.", hostname)

	server.Setup()
}
