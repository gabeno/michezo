package main

import (
	"io"
)

type FileSystemStore struct {
	database io.ReadSeeker
}

func (f *FileSystemStore) GetLeague() []Player {
	f.database.Seek(0, 0)
	league, _ := NewLeague(f.database)
	return league
}

func (f *FileSystemStore) GetPlayerScore(name string) int {
	var win int
	for _, player := range f.GetLeague() {
		if player.Name == name {
			win = player.Wins
			break
		}
	}
	return win
}
