package game

type Series struct {
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	CreatedAt string  `json:"createdAt"`
	Slug      string  `json:"slug"`
	Games     []*Game `json:"games"`
}
