package model

type Game struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	Slug      string `json:"slug"`
}

func (game Game) Save() (Game, error) {
	//conn, _ := rdb.GetConn()
	//
	//_, err := conn.Exec(context.Background(), "INSERT INTO games(slug, name) VALUES($1,$2)", game.Slug, game.Name)
	//
	//if err != nil {
	//	log.Println("Exec error", err)
	//}
	return game, nil
}
