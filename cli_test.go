package poker_test

import (
	"strings"
	"testing"

	"github.com/gabeno/poker/v1"
)

func TestCli(t *testing.T) {
	t.Run("record a win for a chris", func(t *testing.T) {
		in := strings.NewReader("Chris wins/n")
		playerStore := &poker.StubPlayerStore{}

		cli := &poker.Cli{playerStore, in}
		cli.PlayPoker()

		assertPlayerWin(t, playerStore, "Chris")
	})

	t.Run("record a win for a Cleo", func(t *testing.T) {
		in := strings.NewReader("Cleo wins/n")
		playerStore := &poker.StubPlayerStore{}

		cli := &poker.Cli{playerStore, in}
		cli.PlayPoker()

		assertPlayerWin(t, playerStore, "Cleo")
	})
}

func assertPlayerWin(t testing.TB, store *poker.StubPlayerStore, winner string) {
	if len(store.WinCalls) != 1 {
		t.Fatalf("got %d calls to RecordWin want %d", len(store.WinCalls), 1)
	}

	if store.WinCalls[0] != winner {
		t.Errorf("did not record the correct winner, got %s want %s", store.WinCalls[0], winner)
	}
}
