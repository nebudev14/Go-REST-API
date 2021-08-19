package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type album struct {
	ID		string	`json:"id"`
	Title	string	`json:"title"`
	Artist	string	`json:"artist"`
	Price	float64	`json:"price"`	
}

var albums = []album {
	{ID: "1", Title: "His Theme", Artist: "Toby Fox", Price: 4.99},
	{ID: "2", Title: "Introduce a little anarchy", Artist: "Hans Zimmer", Price: 4.99},
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)

	router.Run("localhost:3000")
}
