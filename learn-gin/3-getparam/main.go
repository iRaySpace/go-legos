package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/region/:id", func(c *gin.Context) {
		id := c.Param("id")
		if id == "10" {
			c.JSON(http.StatusOK, gin.H{
				"name":         "Misamis Oriental",
				"capital_city": "Cagayan de Oro",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "Maintenance mode",
		})
	})
	router.Run()
}
