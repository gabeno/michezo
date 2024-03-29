package poker

type InMemoryPlayerStore struct {
	store map[string]int
}

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{map[string]int{}}
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return i.store[name]
}

func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.store[name]++ // candidate for race condition, use mutexes
}

func (i *InMemoryPlayerStore) GetLeague() []Player {
	var league []Player
	for name, wins := range i.store {
		league = append(league, Player{name, wins})
	}
	return league
}
