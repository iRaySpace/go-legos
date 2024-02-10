package main

import (
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	api := e.Group("/api")
	bindTodoApi(api)

	e.Logger.Fatal(e.Start(":8000"))
}
