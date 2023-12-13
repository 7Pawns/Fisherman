package server

import (
	"Fisherman/server/hooks"
)

func Setup() {
	hook := hooks.Hook{}
	hook.StartListen()
}
