package main

import (
	"crud/model"
	"io"
	"text/template"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

var dogMap = model.DogMap{}

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func render(c echo.Context, cmp templ.Component) error {
	return cmp.Render(c.Request().Context(), c.Response())
}

func rowsHandler(c echo.Context) error {
	dogSlice := []model.Dog{}
	for _, dog := range dogMap {
		dogSlice = append(dogSlice, dog)
	}
	d := model.Dog{
		Id:    "1",
		Name:  "bl",
		Breed: "jk",
	}
	return c.Render(200, "public/views/dog_row.html", d)
}
func main() {
	// views.addDog("Pequenito", "Chiuaua")
	// views.addDog("El Mejor", "German Shephard")
	t := &Template{
		templates: template.Must(template.ParseFiles("public/views/dog_row.html")),
	}
	e := echo.New()
	e.Renderer = t
	e.Static("/", "public")
	e.GET("/rows", rowsHandler)
	e.Logger.Fatal(e.Start(":8080"))
}
