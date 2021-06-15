package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type SomeStructExplicit struct {
	Test  int    `json:"test"`
	Butts string `json:"-"`
}

type SomeStructImplicit struct {
	Test  int `json:"test"`
	Butts string
}

func main() {
	e := echo.New()

	e.GET("/explicit", func(c echo.Context) error {
		return c.JSON(http.StatusOK, SomeStructExplicit{Test: 1, Butts: "explicit"})
	})

	e.GET("/implicit", func(c echo.Context) error {
		return c.JSON(http.StatusOK, SomeStructImplicit{Test: 1})
	})

	e.Logger.Fatal(e.Start(":1323"))
}
