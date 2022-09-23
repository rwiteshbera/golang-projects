package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Movie struct {
	ID       string    `json: "id"`
	ISBN     string    `json: "isbn"`
	TITLE    string    `json: "title"`
	DIRECTOR *Director `json: "director"`
}

type Director struct {
	Firstname string `json: "firstname"`
	Lastname  string `json: "lastname"`
}

var movies []Movie

// Get all the movies
func getMovies(c *gin.Context) {
	// encoding movies of type struct as json
	c.JSON(http.StatusOK, movies)
}

func getMovie(c *gin.Context) {
	params := c.Param("id")
	for _, item := range movies {
		if item.ID == params {
			c.JSON(http.StatusOK, item)
			return
		}
	}
}

// Delete a movie
func deleteMovie(c *gin.Context) {
	params := c.Param("id")

	for index, item := range movies {
		if item.ID == params {
			movies = append(movies[:index], movies[index+1:]...)
			c.JSON(http.StatusOK, gin.H{
				"message": movies,
				"params":  params,
			})
			break
		}
	}
}

func createMovie(c *gin.Context) {
	var movie Movie
	c.BindJSON(&movie)
	movie.ID = strconv.Itoa(rand.Intn(1000000000))
	movies = append(movies, movie)
	c.JSON(http.StatusOK, movie)
}

func updateMovie(c *gin.Context) {
	params := c.Param("id")

	for index, item := range movies {
		if item.ID == params {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			c.BindJSON(&movie)
			movie.ID = params
			movies = append(movies, movie)
			c.JSON(http.StatusOK, movies)
		}
	}
}

func main() {
	router := gin.Default()
	router.SetTrustedProxies([]string{"192.168.0.1"})

	router.GET("/movies", getMovies)
	router.GET("/movies/:id", getMovie)
	router.POST("/create", createMovie)
	router.PATCH("/movies/:id", updateMovie)
	router.DELETE("/movies/:id", deleteMovie)

	fmt.Println("Server is listening at PORT 8000")
	router.Run(":8000")
}
