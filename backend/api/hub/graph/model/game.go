package model

type Game struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	Slug      string `json:"slug"`
}
