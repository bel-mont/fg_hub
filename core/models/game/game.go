package game

type Game struct {
	ID         string       `json:"id"`
	Name       string       `json:"name"`
	CreatedAt  string       `json:"createdAt"`
	Slug       string       `json:"slug"`
	SeriesID   string       `json:"seriesId"`
	Characters []*Character `json:"games"`
}
