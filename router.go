package main

import "github.com/gorilla/mux"

func SetupRoutes(router *mux.Router)() {
	router.HandleFunc("/send",SendMessage).Methods("POST")

}