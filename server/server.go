package server

import (
	"Fisherman/server/baits"
)

func Setup() {
	bait := baits.Bait{}
	bait.NewBait()
	bait.StartListen()
}
