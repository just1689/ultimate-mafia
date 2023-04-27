package domain

func NewGame() *Game {
	return &Game{
		Admin: Player{},
		Hand:  make([]*Card, 0),
	}
}

type Game struct {
	Admin Player  `json:"admin"`
	Hand  []*Card `json:"deck"`
}
