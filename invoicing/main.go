package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"irayspace.com/invoicing/app/invoice"
	_ "irayspace.com/invoicing/docs"
)

// @title Invoicing API
// @version 1.0
// @description API for minimial invoicing needs.
func main() {
	e := echo.New()
	e.GET("/swagger/*", echoSwagger.EchoWrapHandler())
	e.GET("/health", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})

	g := e.Group("/api/v1")
	invoice.BindInvoice(g)

	e.Logger.Fatal(e.Start(":8000"))
}
