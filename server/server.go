package server

import (
	"ShyGuy/server/baits"
)

func Setup() {
	bait := baits.Bait{}
	bait.NewBait()
	bait.StartListen()
}
