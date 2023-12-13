package hooks

import (
	"log"
	"net/http"
)

const (
	hookPage      string = "/hook"
	hookPort      string = ":8080"
	hookedMessage string = "Hooked! Something took the bait"
)

type Hook struct {
	server *http.Server
}

func (hook Hook) NewBait() {

}

func sayHello(w http.ResponseWriter, r *http.Request) {
	log.Println(hookedMessage)
}

func (hook Hook) StartListen() {
	http.HandleFunc(hookPage, sayHello)
	log.Fatal(http.ListenAndServe(hookPort, nil))
}
