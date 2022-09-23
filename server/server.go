package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/home", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World!",
		})
	})

	// Router with parameter
	router.GET("/home/:name", func(c *gin.Context) {
		name := c.Param("name")
		// c.JSON(http.StatusOK, gin.H{
		// 	"Name": name,
		// })
		c.String(http.StatusOK, name)
	})

	// Query String
	// /welcome?firstname=Rwitesh&lastname=Bera
	router.GET("/welcome", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "Guest")
		lastname := c.Query("lastname")

		c.JSON(http.StatusOK, gin.H{
			"FirstName": firstname,
			"LastName":  lastname,
		})
	})

	router.Run(":8000")
}
