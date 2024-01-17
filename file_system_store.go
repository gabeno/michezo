package main

import (
	"io"
)

type FileSystemStore struct {
	database io.Reader
}

func (f *FileSystemStore) GetLeague() []Player {
	league, _ := NewLeague(f.database)
	return league
}
