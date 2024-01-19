package poker

type Cli struct {
	store PlayerStore
}

func (c *Cli) PlayPoker() {
	c.store.RecordWin("Cleo")
}
