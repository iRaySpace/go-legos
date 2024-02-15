package main

import (
	"github.com/irayspace/go-legos/learn-echo/6-htmx/internal"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	internal.SetupHandler(e)

	e.Logger.Fatal(e.Start(":8000"))
}
