package main

import (
	"ShyGuy/agent/info"
	"fmt"
)

func main() {
	hostname := info.GetHostname()
	fmt.Printf("ShyGuy initialized on %s.", hostname)
}
