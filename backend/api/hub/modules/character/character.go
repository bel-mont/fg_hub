package character

import (
	"context"
	"fg_hub/backend/api/hub/modules/game"
	"github.com/bel-mont/fg_hub/backend/db/rdb"
	"log"
)

type Character struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	Game      *game.Game
}

//#2
func (character Character) Save() string {
	//#3
	//stmt, err := rdb.Conn.Prepare(context.Background(), "CharacterInsert", "INSERT INTO characters(name, game_id) VALUES(?,?)")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//#4
	err := rdb.Conn.QueryRow(context.Background(), " INSERT INTO characters(name, game_id) VALUES($1,$2)", character.Name, "c2d29867-3d0b-d497-9191-18a9d8ee7812")
	//res, err := stmt.Exec(character.Title, character.Address)
	if err != nil {
		log.Fatal("Panic!!!!", err)
	}
	//#5
	//id, err := res.()
	//if err != nil {
	//	log.Fatal("Error:", err.Error())
	//}
	log.Print("Row inserted!")
	return "sample"
}
