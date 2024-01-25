package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		region := c.Query("region")
		if region == "10" {
			c.JSON(http.StatusOK, gin.H{
				"name":         "Misamis Oriental",
				"founded":      1939,
				"capital_city": "Cagayan de Oro",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "Limited use only",
		})
	})
	router.Run()
}
