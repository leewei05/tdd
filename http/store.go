package main

import (
	"io"
)

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{map[string]int{}}
}

type InMemoryPlayerStore struct {
	store map[string]int
}

func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.store[name]++
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return i.store[name]
}

func (i *InMemoryPlayerStore) GetLeague() []Player {
	var league []Player
	for name, wins := range i.store {
		league = append(league, Player{name, wins})
	}

	return league
}

type FileSystemPlayerStore struct {
	db io.ReadSeeker
}

func (f *FileSystemPlayerStore) GetLeague() []Player {
	f.db.Seek(0, 0)
	league, _ := NewLeague(f.db)

	return league
}
