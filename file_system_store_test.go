package poker_test

import (
	"testing"

	"github.com/gabeno/poker/v1"
)

func TestFileSystemStore(t *testing.T) {

	t.Run("sorted league from a reader", func(t *testing.T) {
		database, cleanDatabase := poker.CreateTempFile(t, `[{"Name": "Cleo", "Wins": 10},{"Name": "Chris", "Wins": 33}]`)
		defer cleanDatabase()
		store, err := poker.NewFileSystemStore(database)

		poker.AssertNoError(t, err)

		got := store.GetLeague()
		want := []poker.Player{
			{"Chris", 33},
			{"Cleo", 10},
		}

		poker.AssertLeague(t, got, want)

		// read again
		got = store.GetLeague()
		poker.AssertLeague(t, got, want)
	})

	t.Run("get player score", func(t *testing.T) {
		database, cleanDatabase := poker.CreateTempFile(t, `[{"Name": "Cleo", "Wins": 10},{"Name": "Chris", "Wins": 33}]`)
		defer cleanDatabase()
		store, err := poker.NewFileSystemStore(database)

		poker.AssertNoError(t, err)

		got := store.GetPlayerScore("Chris")
		want := 33

		poker.AssertScoreEquals(t, got, want)
	})

	t.Run("record win for existing player", func(t *testing.T) {
		database, cleanDatabase := poker.CreateTempFile(t, `[{"Name": "Cleo", "Wins": 10},{"Name": "Chris", "Wins": 33}]`)
		defer cleanDatabase()
		store, err := poker.NewFileSystemStore(database)

		poker.AssertNoError(t, err)

		store.RecordWin("Cleo")

		got := store.GetPlayerScore("Cleo")
		want := 11

		poker.AssertScoreEquals(t, got, want)
	})

	t.Run("record win for new player", func(t *testing.T) {
		database, cleanDatabase := poker.CreateTempFile(t, `[{"Name": "Cleo", "Wins": 10},{"Name": "Chris", "Wins": 33}]`)
		defer cleanDatabase()
		store, err := poker.NewFileSystemStore(database)

		poker.AssertNoError(t, err)

		store.RecordWin("Pepper")

		got := store.GetPlayerScore("Pepper")
		want := 1

		poker.AssertScoreEquals(t, got, want)
	})

	t.Run("works with an empty file", func(t *testing.T) {
		file, clean := poker.CreateTempFile(t, "")
		defer clean()
		_, err := poker.NewFileSystemStore(file)

		poker.AssertNoError(t, err)
	})
}
