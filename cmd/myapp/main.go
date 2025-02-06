package main

import (
	"fmt"
	"htmx-go-demo/internal/config"
	"htmx-go-demo/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

var dogMap = model.DogMap{}
var selectedId = ""

func main() {
	config, err := config.Load()
	if err != nil {
		panic(err)
	}
	fmt.Println(config)
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World")
	})
	e.Logger.Fatal(e.Start(":8080"))
}
