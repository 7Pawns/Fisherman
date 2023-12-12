package baits

import (
	"log"
	"net/http"
)

type Bait struct {
	server *http.Server
}

func (bait Bait) NewBait() {

}

func sayHello(w http.ResponseWriter, r *http.Request) {
	log.Println("Hello Bait!\n")
}

func (bait Bait) StartListen() {
	http.HandleFunc("/bait", sayHello)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
