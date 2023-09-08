package main

import (
	"log"
	"net/http"
)

type InMemoryPlayerStore struct{}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return 123
}

func main() {

	server := &PlayerServer{&InMemoryPlayerStore{}}

	if err := http.ListenAndServe(":5011", server); err != nil {
		log.Fatalf("could not listen on port 5011 %v", err)
	}
}
