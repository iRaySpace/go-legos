package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/hephep", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "hooray"})
	})
	router.Run()
}
