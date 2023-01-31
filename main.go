package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist *Artist `json:"artist"`
	Price  float64 `json:"price"`
}

type Artist struct {
	Firstname string `json:"firstname"`
	LastName  string `json:"lastname"`
}

var Albums = []album{
	{ID: "1", Title: "See Me", Artist: &Artist{Firstname: "Rich", LastName: "Brian"}, Price: 100},
	{ID: "2", Title: "Every Summer Time", Artist: &Artist{Firstname: "Niki", LastName: ""}, Price: 110},
	{ID: "3", Title: "California", Artist: &Artist{Firstname: "Rich", LastName: "Brian"}, Price: 80},
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, Albums)
}

func postAlbums(c *gin.Context) {
	var newAlbum album

	//Call BindJSON to bind the received JSON to newAlbum
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	//add new album to slice
	Albums = append(Albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	//Loop over the list of album, looking for the match id
	for _, item := range Albums {
		if item.ID == id {
			c.IndentedJSON(http.StatusOK, item)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func updateAlbumByID(c *gin.Context) {
	id := c.Param("id")

	for index, item := range Albums {
		if item.ID == id {

			//append new data to macthed id data
			Albums = append(Albums[:index], Albums[index+1:]...)

			var newAlbum album

			//Call BindJSON to bind the received JSON to newAlbum
			if err := c.BindJSON(&newAlbum); err != nil {
				return
			}

			//add new album to slice
			Albums = append(Albums, newAlbum)
			c.IndentedJSON(http.StatusCreated, newAlbum)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func deleteAlbumById(c *gin.Context) {
	id := c.Param("id")

	for index, item := range Albums {
		if item.ID == id {
			Albums = append(Albums[:index], Albums[index+1:]...)
			break
		}
	}
	c.IndentedJSON(http.StatusOK, Albums)
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Album successfully deleted"})
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.PUT("/albums/:id", updateAlbumByID)
	router.DELETE("/albums/:id", deleteAlbumById)
	router.POST("/albums", postAlbums)

	router.Run("localhost: 3030")
}
