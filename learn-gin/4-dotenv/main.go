package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	buildNo := os.Getenv("BUILD_NO")

	router := gin.Default()
	router.GET("/about", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"buildNo": buildNo,
		})
	})
	router.Run()
}
