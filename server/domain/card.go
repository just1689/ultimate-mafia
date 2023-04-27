package domain

type Card struct {
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Color       string     `json:"color"`
	Abilities   []*Ability `json:"abilities"`
}
