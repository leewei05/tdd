package main

import (
	"log"
	"net/http"
)

type InMemoryPlayerStore struct{}

func main() {
	playerStore := InMemoryPlayerStore{}
	server := &PlayerServer{&playerStore}
	log.Fatal(http.ListenAndServe(":8000", server))
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {

	return 123
}
