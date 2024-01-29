package main

import (
	"os"
	"fmt"
	"time"
	"strings"
	"net/http"
	"github.com/gin-gonic/gin"
	"aidanwoods.dev/go-paseto"
	"github.com/joho/godotenv"
)

type LoginDto struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	r := gin.Default()
	r.POST("/login", func(c *gin.Context) {
		var loginDto LoginDto
		if err := c.BindJSON(&loginDto); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		if loginDto.Username != "admin" || loginDto.Password != "admin" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})
			return
		}

		token := paseto.NewToken()
		now := time.Now()
		token.SetIssuedAt(now)
		token.SetExpiration(now.Add(30 * time.Second))
		token.Set("username", loginDto.Username)

		key, err := paseto.V4SymmetricKeyFromHex(os.Getenv("AUTH_KEY"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}
		encrypted := token.V4Encrypt(key, nil)

		c.JSON(http.StatusOK, gin.H{
			"token": encrypted,
		})
	})

	r.GET("/protected", func(c *gin.Context) {
		bearerToken := c.GetHeader("Authorization")
		if bearerToken == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})
			return
		}
		
		token := strings.TrimPrefix(bearerToken, "Bearer ")

		key, err := paseto.V4SymmetricKeyFromHex(os.Getenv("AUTH_KEY"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		parser := paseto.NewParser()
		parsedToken, err := parser.ParseV4Local(key, token, nil)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": err.Error(),
			})
			return
		}

		username, err := parsedToken.GetString("username")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Hello, %s!", username),
		})
	})

	r.Run()
}