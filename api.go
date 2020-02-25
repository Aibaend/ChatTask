package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type ChanellResponse struct {
	Type      string   `json:"type"`
	ChanellID string   `json:"chanell_id"`
	Message   *Message `json:"message,omitempty"`
}

type Message struct {
	To      string `json:"to"`
	Name    string `json:"name"`
	Message string `json:"message"`
}

var Response = func(w http.ResponseWriter, r *http.Request) ChanellResponse {
	log.Println(r.RequestURI, r.Method)
	channel := ChanellResponse{}
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(&channel)
	if err != nil {
		log.Println(r.RequestURI, r.Method, err.Error())
	}
	return channel
}

var SendMessage = func(w http.ResponseWriter, r *http.Request, hub *Hub) {
	log.Println(r.RequestURI, r.Method)
	channel := ChanellResponse{}
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(&channel)
	if err != nil {
		log.Println(r.RequestURI, r.Method, err.Error())
	}
	if channel.Type == "event" {
		hub.broadcast <- []byte(channel.Message.Message)
	}

}

func (manager *Hub) send(message []byte, ignore *Client) {
	for conn := range manager.clients {
		if conn != ignore {
			conn.send <- message
		}
	}
}
