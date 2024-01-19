package poker

import (
	"bufio"
	"io"
	"strings"
)

type Cli struct {
	store PlayerStore
	in    *bufio.Scanner
}

func NewCli(store PlayerStore, in io.Reader) *Cli {
	return &Cli{
		store: store,
		in:    bufio.NewScanner(in),
	}
}

func (c *Cli) PlayPoker() {
	userInput := c.ReadLine()
	c.store.RecordWin(extractPlayerName(userInput))
}

func extractPlayerName(text string) string {
	return strings.Replace(text, " wins/n", "", 1)
}

func (c *Cli) ReadLine() string {
	c.in.Scan()
	return c.in.Text()
}
