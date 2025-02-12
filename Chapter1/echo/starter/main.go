package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.Static("/", "public")
	e.GET("/version", func(c echo.Context) error {
		return c.String(http.StatusOK, "this is the version")
	})
	e.Logger.Fatal(e.Start(":8080"))
}
