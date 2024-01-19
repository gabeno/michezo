package poker

import "testing"

func TestCli(t *testing.T) {
	t.Run("record a win for a player", func(t *testing.T) {
		playerStore := &StubPlayerStore{}

		cli := &Cli{playerStore}
		cli.PlayPoker()

		if len(playerStore.winCalls) != 1 {
			t.Fatal("expected a win call but did not get any")
		}
	})
}
