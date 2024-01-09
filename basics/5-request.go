package main

import (
	"net/http"
	"strings"
	"os"
)

// USER_ID=1234567890 DISCORD_URL=https://discord.com/api/webhooks/1234567890/abcdefghijklmnopqrstuvwxyz go run 5-request.go
func main() {
	discordUrl := os.Getenv("DISCORD_URL")
	userId := os.Getenv("USER_ID")
	body := `{"content": "Hello, <@` + userId + `>!"}`
	_, err := http.Post(discordUrl, "application/json", strings.NewReader(body))
	if err != nil {
		panic(err)
	}
}