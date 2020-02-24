package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type ChanellResponse struct {
	Type      string   `json:"type"`
	ChanellID uint     `json:"chanell_id"`
	Message   *Message `json:"message,omitempty"`
}

type Message struct {
	To      string `json:"to"`
	Name    string `json:"name"`
	Message string `json:"message"`
}

var Switcher = func(w http.ResponseWriter, r *http.Request) ChanellResponse {
	log.Println(r.RequestURI, r.Method)
	channel := ChanellResponse{}
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(&channel)
	if err != nil {
		log.Println(r.RequestURI, r.Method, err.Error())
	}
	return channel
}

var SendMessage = func(w http.ResponseWriter, r http.Request) {
	log.Println(r.RequestURI, r.Method)
	channel := ChanellResponse{}
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(&channel)
	if err != nil {
		log.Println(r.RequestURI, r.Method, err.Error())
	}

}
