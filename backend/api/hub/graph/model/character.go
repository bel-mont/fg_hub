package model

import (
	"log"
)

type Character struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	Game      *Game
}

//#2
func (character Character) Save() Character {
	//err := rdb.Conn.QueryRow(context.Background(), " INSERT INTO characters(name, game_id) VALUES($1,$2)", character.Name, "c2d29867-3d0b-d497-9191-18a9d8ee7812")
	////res, err := stmt.Exec(character.Title, character.Address)
	//if err != nil {
	//	log.Fatal("Panic!!!!", err)
	//}
	log.Print("Row inserted!")
	return character
}
