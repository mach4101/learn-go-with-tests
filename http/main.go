package main

import (
	"log"
	"net/http"
)

func main() {
	handler := http.HandlerFunc(PlayerServer)

	if err := http.ListenAndServe(":5011", handler); err != nil {
		log.Fatalf("could not listen on port 5011 %v", err)
	}
}
