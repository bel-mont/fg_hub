package game

type Character struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"createdAt"`
	GameID    string `json:"gameId"`
	Slug      string `json:"slug"`
}
