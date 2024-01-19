package poker

import (
	"strings"
	"testing"
)

func TestCli(t *testing.T) {
	t.Run("record a win for a player", func(t *testing.T) {
		in := strings.NewReader("Chris /n")
		playerStore := &StubPlayerStore{}

		cli := &Cli{playerStore, in}
		cli.PlayPoker()

		if len(playerStore.winCalls) != 1 {
			t.Fatal("expected a win call but did not get any")
		}

		got := playerStore.winCalls[0]
		want := "Chris"

		if got != want {
			t.Errorf("did not record the correct winner, got %s want %s", got, want)
		}
	})
}
