package main

import (
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.Debug = true
	e.Renderer = InitTemplateRenderer()

	p := Pages{}
	e.GET("/", p.index)

	e.Logger.Fatal(e.Start(":8000"))
}
