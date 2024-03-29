package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Создаем альбом
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

type error struct {
	Error string `json:"error"`
}

// Заполняем альбом
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltraine", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mullingan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(c *gin.Context) {
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		c.IndentedJSON(http.StatusBadRequest, error{"bad_request"})
		return
	}
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumById(c *gin.Context) {
	id := c.Param("id")
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"massage": "album not found"})
}

func deleteAlbumById(c *gin.Context) {
	id := c.Param("id")

	for i, a := range albums {
		if a.ID == id {
			albums = append(albums[:i], albums[i+1:]...)
			c.IndentedJSON(http.StatusNoContent, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"massage": "album not found"})
}

func updateAlbumById(c *gin.Context) {
	id := c.Param("id")

	for i := range albums {
		if albums[i].ID == id {
			c.BindJSON(&albums[i])
			c.IndentedJSON(http.StatusOK, albums[i])
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"massage": "album not found"})
}

func getRouter() *gin.Engine { // gin.Engine - это основной компонент фреймворка Gin, который представляет собой маршрутизатор (router) для обработки HTTP-запросов и управления маршрутами (routes) веб-приложения.
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumById)
	router.DELETE("/albums/:id", deleteAlbumById)
	router.PUT("/albums/:id", updateAlbumById)
	router.POST("/albums", postAlbums)
	return router
}

func main() {
	router := getRouter()
	router.Run("localhost:8080")
}
