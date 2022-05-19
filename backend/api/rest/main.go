package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type game struct {
	Id    string `json:"id"`
	Title string `json:"title"`
}

// games slice to seed record game data.
var games = []game{
	{Id: "1", Title: "Street Fighter 5"},
	{Id: "2", Title: "Granblue Fantasy Versus"},
	{Id: "3", Title: "Guilty Gear Strive"},
}

// getgames responds with the list of all games as JSON.
func getGames(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, games)
}

// postGames adds an game from JSON received in the request body.
func postGames(c *gin.Context) {
	var newgame game

	// Call BindJSON to bind the received JSON to
	// newgame.
	if err := c.BindJSON(&newgame); err != nil {
		return
	}

	// Add the new game to the slice.
	games = append(games, newgame)
	c.IndentedJSON(http.StatusCreated, newgame)
}

// getGameById locates the game whose ID value matches the id
// parameter sent by the client, then returns that game as a response.
func getGameById(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of games, looking for
	// an game whose ID value matches the parameter.
	for _, a := range games {
		if a.Id == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "game not found"})
}

func getRdbData(c *gin.Context) {
	// rdbTest := rdb.ReadTest()
	// c.IndentedJSON(http.StatusOK, rdbTest)
}

func test() {
	// rdb.Test()
	router := gin.Default()
	router.GET("/game", getGames)
	router.GET("/game/:id", getGameById)
	router.GET("/sf5", getRdbData)
	router.POST("/game", postGames)
	router.Run("localhost:8080")
}
