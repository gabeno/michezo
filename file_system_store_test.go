package main

import (
	"os"
	"testing"
)

func TestFileSystemStore(t *testing.T) {

	t.Run("league from a reader", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[{"Name": "Cleo", "Wins": 10},{"Name": "Chris", "Wins": 33}]`)
		defer cleanDatabase()
		store, err := NewFileSystemStore(database)

		assertNoError(t, err)

		got := store.league
		want := []Player{
			{"Cleo", 10},
			{"Chris", 33},
		}

		assertLeague(t, got, want)

		// read again
		got = store.GetLeague()
		assertLeague(t, got, want)
	})

	t.Run("get player score", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[{"Name": "Cleo", "Wins": 10},{"Name": "Chris", "Wins": 33}]`)
		defer cleanDatabase()
		store, err := NewFileSystemStore(database)

		assertNoError(t, err)

		got := store.GetPlayerScore("Chris")
		want := 33

		assertScoreEquals(t, got, want)
	})

	t.Run("record win for existing player", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[{"Name": "Cleo", "Wins": 10},{"Name": "Chris", "Wins": 33}]`)
		defer cleanDatabase()
		store, err := NewFileSystemStore(database)

		assertNoError(t, err)

		store.RecordWin("Cleo")

		got := store.GetPlayerScore("Cleo")
		want := 11

		assertScoreEquals(t, got, want)
	})

	t.Run("record win for new player", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[{"Name": "Cleo", "Wins": 10},{"Name": "Chris", "Wins": 33}]`)
		defer cleanDatabase()
		store, err := NewFileSystemStore(database)

		assertNoError(t, err)

		store.RecordWin("Pepper")

		got := store.GetPlayerScore("Pepper")
		want := 1

		assertScoreEquals(t, got, want)
	})

	t.Run("works with an empty file", func(t *testing.T) {
		file, clean := createTempFile(t, "")
		defer clean()
		_, err := NewFileSystemStore(file)

		assertNoError(t, err)
	})
}

func assertScoreEquals(t testing.TB, got, want int) {
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func assertNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("didn't expect an error but got one, %v", err)
	}
}

func createTempFile(t testing.TB, data string) (*os.File, func()) {
	tempFile, err := os.CreateTemp("", "db")

	if err != nil {
		t.Fatalf("could not create temp file %v", err)
	}

	tempFile.Write([]byte(data))

	removeFile := func() {
		tempFile.Close()
		os.Remove(tempFile.Name())
	}

	return tempFile, removeFile
}
