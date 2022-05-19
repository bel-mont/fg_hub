package game

type Character struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	GameID    string `json:"game_id"`
	Slug      string `json:"slug"`
}
