package domain

import "github.com/oklog/ulid/v2"

func NewGame() *Game {
	return &Game{
		ID:         ulid.Make().String(),
		JoinSecret: ulid.Make().String(),
		Admin:      Player{},
		Hand:       make([]*Card, 0),
	}
}

type Game struct {
	ID         string  `json:"id"`
	JoinSecret string  `json:"joinSecret"`
	Admin      Player  `json:"admin"`
	Hand       []*Card `json:"deck"`
}
