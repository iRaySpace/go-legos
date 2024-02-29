package main

import (
	"github.com/jackc/pgtype"
	"github.com/labstack/echo/v4"
)

type TimeRequest struct {
	Time string `json:"time"`
}

type TimeResponse struct {
	Hours   int64 `json:"hours"`
	Minutes int64 `json:"minutes"`
	Seconds int64 `json:"seconds"`
}

func main() {
	e := echo.New()
	e.POST("/", func(c echo.Context) error {
		t := new(TimeRequest)
		if err := c.Bind(t); err != nil {
			return c.JSON(400, map[string]interface{}{
				"message": "invalid request data",
			})
		}

		pt := new(pgtype.Time)
		err := pt.DecodeText(nil, []byte(t.Time)) // refactor this into a util function
		if err != nil {
			return c.JSON(400, map[string]interface{}{
				"message": "unable to parse data",
			})
		}

		seconds := pt.Microseconds / 1000000
		res := TimeResponse{
			Hours:   (seconds / 60 / 60 % 24),
			Minutes: (seconds / 60 % 60),
			Seconds: (seconds % 60),
		}

		return c.JSON(200, map[string]interface{}{
			"data": res,
		})
	})
	e.Logger.Fatal(e.Start(":8000"))
}
