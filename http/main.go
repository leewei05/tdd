package main

import (
	"log"
	"net/http"
)

func main() {
	playerStore := NewInMemoryPlayerStore()
	server := &PlayerServer{playerStore}
	log.Fatal(http.ListenAndServe(":8000", server))
}
