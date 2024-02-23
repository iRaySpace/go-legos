package main

import (
	"context"
	"log"

	"github.com/irayspace/go-legos/learn-echo/7-sqlc/data"
	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v4"
)

// import "github.com/labstack/echo/v4"

func main() {
	e := echo.New()

	ctx := context.Background()
	conn, err := pgx.Connect(ctx, "postgres://postgres:123@localhost:5432/postgres")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(ctx)
	queries := data.New(conn)

	e.GET("/employees", func(c echo.Context) error {
		employees, err := queries.ListEmployees(ctx)
		if err != nil {
			return c.JSON(400, map[string]interface{}{
				"statusCode": 400,
				"message":    "unable to fetch",
			})
		}

		return c.JSON(200, map[string]interface{}{
			"data":  employees,
			"total": len(employees),
		})
	})
	e.POST("/employees", func(c echo.Context) error {
		employee := new(data.CreateEmployeeParams)
		if err := c.Bind(employee); err != nil {
			return c.JSON(400, map[string]interface{}{"statusCode": 400, "message": "unable to create"})
		}

		newEmployee, err := queries.CreateEmployee(ctx, *employee)
		if err != nil {
			return c.JSON(400, map[string]interface{}{"statusCode": 400, "message": "unable to create"})
		}

		return c.JSON(201, map[string]interface{}{
			"data": newEmployee,
		})
	})

	e.Logger.Fatal(e.Start(":8000"))
}
