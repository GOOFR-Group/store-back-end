package storage

type StateGame string

const (
	StateGameActive   StateGame = "active"
	StateGameInactive StateGame = "inactive"
	StateGameUpcoming StateGame = "upcoming"
)
