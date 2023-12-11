package info

import (
	"fmt"
	"os"
)

func GetHostname() string {
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return hostname
}
