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

func postAlbum(c *gin.Context) {
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbum)

	router.Run("localhost:3000")
}
