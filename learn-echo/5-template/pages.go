package main

import (
	"github.com/labstack/echo/v4"
)

type Pages struct{}

func (p *Pages) index(c echo.Context) error {
	return c.Render(200, "index.html", map[string]interface{}{
		"title": "Hello, world!",
	})
}
