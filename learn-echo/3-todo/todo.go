package main

import (
	"strconv"

	"github.com/labstack/echo/v4"
)

func bindTodoApi(g *echo.Group) {
	api := todoApi{}
	g.POST("/todos", api.create)
	g.GET("/todos", api.list)
	g.PUT("/todos/:id", api.update)
	g.DELETE("/todos/:id", api.delete)
}

type todo struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type todoApi struct {
	currentID int
	todos     []todo
}

func (api *todoApi) create(c echo.Context) error {
	t := new(todo)
	if err := c.Bind(t); err != nil {
		return c.JSON(400, map[string]string{"message": "Invalid request"})
	}

	t.ID = api.currentID
	api.currentID = api.currentID + 1
	api.todos = append(api.todos, *t)

	return c.JSON(201, t)
}

func (api *todoApi) list(c echo.Context) error {
	if api.todos == nil {
		api.todos = []todo{}
	}
	return c.JSON(200, api.todos)
}

func (api *todoApi) update(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(400, map[string]string{"message": "Invalid request"})
	}

	updatedT := new(todo)
	if err := c.Bind(updatedT); err != nil {
		return c.JSON(400, map[string]string{"message": "Invalid request"})
	}

	updatedI := -1
	for i, t := range api.todos {
		if t.ID == id {
			updatedI = i
		}
	}

	if updatedI == -1 {
		return c.JSON(400, map[string]string{"message": "Invalid request"})
	}

	api.todos[updatedI].Name = updatedT.Name
	return c.JSON(200, api.todos[updatedI])
}

func (api *todoApi) delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(400, map[string]string{"message": "Invalid request"})
	}

	var filteredT []todo
	var hasId bool
	for _, t := range api.todos {
		if t.ID != id {
			filteredT = append(filteredT, t)
		}
		if t.ID == id {
			hasId = true
		}
	}

	if !hasId {
		return c.JSON(404, map[string]string{"message": "Todo does not exists"})
	}

	api.todos = filteredT

	return c.JSON(200, map[string]string{"message": "Todo deleted"})
}
