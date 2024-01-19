package poker

import (
	"bufio"
	"io"
	"strings"
)

type Cli struct {
	store PlayerStore
	in    io.Reader
}

func (c *Cli) PlayPoker() {
	reader := bufio.NewScanner(c.in)
	reader.Scan()
	c.store.RecordWin(extractPlayerName(reader.Text()))
}

func extractPlayerName(text string) string {
	return strings.Replace(text, " wins/n", "", 1)
}
