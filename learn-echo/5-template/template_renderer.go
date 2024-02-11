package main

import (
	"io"
	"log"
	"text/template"

	"github.com/labstack/echo/v4"
)

type TemplateRenderer struct {
	*template.Template
}

func (t TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.ExecuteTemplate(w, name, data)
}

func InitTemplateRenderer() TemplateRenderer {
	t := template.New("")
	t, err := t.ParseGlob("templates/*.html")
	if err != nil {
		log.Fatal(err)
	}
	return TemplateRenderer{t}
}
