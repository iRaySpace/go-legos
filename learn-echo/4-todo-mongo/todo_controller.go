package main

import (
	"github.com/labstack/echo/v4"
)

type TodoController struct {
	service *TodoService
}

func bindTodo(g *echo.Group) {
	c := TodoController{service: &TodoService{}}
	g.POST("/todos", c.create)
	g.GET("/todos", c.list)
	g.PUT("/todos/:id", c.update)
	g.DELETE("/todos/:id", c.delete)
}

func (tc *TodoController) create(c echo.Context) error {
	data := new(Todo)
	if err := c.Bind(data); err != nil {
		return c.JSON(400, map[string]string{"message": "Invalid request"})
	}
	return c.JSON(201, tc.service.Create(data))
}

func (tc *TodoController) list(c echo.Context) error {
	return c.JSON(200, tc.service.GetAll())
}

func (tc *TodoController) update(c echo.Context) error {
	updatedT := new(Todo)
	if err := c.Bind(updatedT); err != nil {
		return c.JSON(400, map[string]string{"message": "Invalid request"})
	}

	if err := tc.service.Update(c.Param("id"), updatedT); err != nil {
		return c.JSON(404, map[string]string{"message": "Unable to update, document has not been found"})
	}

	return c.JSON(200, map[string]string{"message": "Document has been updated"})
}

func (tc *TodoController) delete(c echo.Context) error {
	err := tc.service.Delete(c.Param("id"))
	if err != nil {
		return c.JSON(404, map[string]string{"message": "Unable to delete, no document has found"})
	}
	return c.JSON(200, map[string]string{"message": "Successfully deleted"})
}
