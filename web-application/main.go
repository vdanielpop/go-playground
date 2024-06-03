package main

import (
	"log"
	"net/http"
)

type InMemoryPlayerStore struct {
	store map[string]int
}

func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.store[name]++
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return i.store[name]
}

func main() {
	server := &PlayerServer{&InMemoryPlayerStore{map[string]int{}}}
	log.Fatal(http.ListenAndServe(":5000", server))
}
