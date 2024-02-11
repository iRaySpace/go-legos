package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	e := echo.New()

	g := e.Group("/api")
	bindTodo(g)

	InitMongo()
	defer DisconnectMongo()

	e.Logger.Fatal(e.Start(":8000"))
}
