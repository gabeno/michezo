package main

import (
	"encoding/json"
	"io"
)

type FileSystemStore struct {
	database io.Writer
	league   League
}

func NewFileSystemStore(database io.ReadWriteSeeker) *FileSystemStore {
	database.Seek(0, 0)
	league, _ := NewLeague(database)
	return &FileSystemStore{
		database: &tape{database},
		league:   league,
	}
}

func (f *FileSystemStore) GetLeague() League {
	return f.league
}

func (f *FileSystemStore) GetPlayerScore(name string) int {
	player := f.league.Find(name)

	if player != nil {
		return player.Wins
	}

	return 0
}

func (f *FileSystemStore) RecordWin(name string) {
	player := f.league.Find(name)

	if player != nil {
		player.Wins++
	} else {
		f.league = append(f.league, Player{name, 1})
	}

	json.NewEncoder(f.database).Encode(&f.league)
}
