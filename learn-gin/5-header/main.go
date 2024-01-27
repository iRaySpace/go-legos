package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/restricted", func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "Bearer" {
			c.JSON(http.StatusOK, gin.H{
				"message": "You are good!",
			})
			return
		}
		c.JSON(http.StatusForbidden, gin.H{
			"message": "User not allowed",
		})
	})
	r.Run()
}
