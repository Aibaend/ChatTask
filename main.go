package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var addr = flag.String("addr", ":8080", "http service address")

func serveHome(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "home.html")
}

func main() {
	flag.Parse()
	router := mux.NewRouter()
	hub := newHub()
	go hub.run()
	router.HandleFunc("/", serveHome)
	router.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})

	router.HandleFunc("/send/message", func(writer http.ResponseWriter, request *http.Request) {
		SendMessage(writer, request, hub)
	})
	err := http.ListenAndServe(*addr, router)
	fmt.Print("Listening on port:8080")
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
