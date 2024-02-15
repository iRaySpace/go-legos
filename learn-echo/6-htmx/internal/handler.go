package internal

import (
	"github.com/a-h/templ"
	"github.com/irayspace/go-legos/learn-echo/6-htmx/domain"
	"github.com/irayspace/go-legos/learn-echo/6-htmx/views/admin"
	"github.com/irayspace/go-legos/learn-echo/6-htmx/views/admin/components"
	"github.com/labstack/echo/v4"
)

func SetupHandler(e *echo.Echo) {
	g := e.Group("/admin")
	g.GET("", getIndex)

	e.GET("/data", getData)
}

func getIndex(c echo.Context) error {
	return renderView(c, admin.Index())
}

func getData(c echo.Context) error {
	i := domain.Item{
		Name:  c.QueryParam("item"),
		Price: 100,
	}
	return components.ItemComponent(i).Render(c.Request().Context(), c.Response().Writer)
}

func renderView(c echo.Context, cmp templ.Component) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
	return cmp.Render(c.Request().Context(), c.Response().Writer)
}
