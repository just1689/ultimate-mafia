package domain

import "github.com/google/uuid"

func NewPlayer(title string) *Player {
	return &Player{
		ID:     uuid.New().String(),
		Secret: uuid.New().String(),
		Title:  title,
		Alive:  true,
		Card:   nil,
	}
}

func NewAdmin(title string) *Player {
	r := NewPlayer(title)
	r.Alive = false
	return r
}

type Player struct {
	ID     string `json:"id"`
	Secret string `json:"secret"`
	Title  string `json:"title"`
	Alive  bool   `json:"alive"`
	Card   *Card  `json:"card"`
}
