package poker

import (
	"strings"
	"testing"
)

func TestCli(t *testing.T) {
	t.Run("record a win for a chris", func(t *testing.T) {
		in := strings.NewReader("Chris wins/n")
		playerStore := &StubPlayerStore{}

		cli := &Cli{playerStore, in}
		cli.PlayPoker()

		assertPlayerWin(t, playerStore, "Chris")
	})

	t.Run("record a win for a Cleo", func(t *testing.T) {
		in := strings.NewReader("Cleo wins/n")
		playerStore := &StubPlayerStore{}

		cli := &Cli{playerStore, in}
		cli.PlayPoker()

		assertPlayerWin(t, playerStore, "Cleo")
	})
}

func assertPlayerWin(t testing.TB, store *StubPlayerStore, winner string) {
	if len(store.winCalls) != 1 {
		t.Fatalf("got %d calls to RecordWin want %d", len(store.winCalls), 1)
	}

	if store.winCalls[0] != winner {
		t.Errorf("did not record the correct winner, got %s want %s", store.winCalls[0], winner)
	}
}
