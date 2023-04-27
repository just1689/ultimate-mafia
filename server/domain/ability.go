package domain

type Ability struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Reusable    bool   `json:"reusable"`
}
