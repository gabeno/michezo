package poker

import "io"

type Cli struct {
	store PlayerStore
	in    io.Reader
}

func (c *Cli) PlayPoker() {
	c.store.RecordWin("Chris")
}
