package main

import (
	"fmt"
	"log"
	"net/http"
	"slices"
	"strings"
	"text/template"

	"github.com/fadykuzman/htmx-pql/model"
	persist "github.com/fadykuzman/htmx-pql/persistence"
	"github.com/fadykuzman/htmx-pql/repositories"
)

var dogRepo *repositories.DogRepository

type dogData struct {
	Dog   model.Dog
	Attrs map[string]string
}

func getDogs(w http.ResponseWriter, r *http.Request) {
	dogs := make([]dogData, 0)

	dogsSlice, err := dogRepo.GetDogs()

	if err != nil {
		log.Fatal(err)
	}

	slices.SortStableFunc(dogsSlice, func(a, b model.Dog) int {
		return strings.Compare(strings.ToLower(a.Name), strings.ToLower(b.Name))
	})

	attrs := make(map[string]string, 0)

	for _, dog := range dogsSlice {
		dd := dogData{
			Dog:   dog,
			Attrs: attrs,
		}
		dogs = append(dogs, dd)
	}

	itemTemplate, err := parseDog()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error parsing dog template: %v", err)
	}

	listTmpl := template.Must(itemTemplate.Clone())

	_, err = listTmpl.ParseFiles("public/dogs.html")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error parsing dogs template: %v", err)
	}
	listTmpl.ExecuteTemplate(w, "dog-rows", dogs)
}

func parseDog() (*template.Template, error) {
	return template.ParseFiles("public/dog.html")
}

func createDog(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	breed := r.FormValue("breed")
	dog := dogRepo.CreateDog(name, breed)
	slice := []model.Dog{
		dog,
	}
	tmpl := template.Must(template.ParseFiles("public/dogs.html"))
	tmpl.Execute(w, slice)
}

func deleteDog(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	dogRepo.DeleteDog(id)
}

var selected_id string

type FormData struct {
	SelectedId  string
	SelectedDog model.Dog
	Attrs       map[string]string
}

func getForm(w http.ResponseWriter, r *http.Request) {
	attrs := map[string]string{
		"hx-on:htmx:after-request": "this.reset()",
	}
	if selected_id != "" {
		attrs["hx-put"] = "/dog" + selected_id
	} else {
		attrs["hx-post"] = "/dog"
		attrs["hx-target"] = "tbody"
		attrs["hx-swap"] = "afterbegin"
	}
	selected_dog := dogRepo.GetDog(selected_id)
	fmt.Printf("dog: %s\n", selected_dog)
	formData := FormData{
		SelectedId:  selected_id,
		SelectedDog: selected_dog,
		Attrs:       attrs,
	}
	tmpl := template.Must(template.ParseFiles("public/form.html"))
	tmpl.Execute(w, formData)

}

func selectDog(w http.ResponseWriter, r *http.Request) {
	selected_id = r.PathValue("id")
	w.Header().Set("HX-Trigger", "selection-change")
}

func main() {
	err := persist.Connect()
	if err != nil {
		log.Println(err)
	}

	defer persist.GetDB().Close()

	dogRepo = repositories.NewDogRepository()
	srvr := http.NewServeMux()

	srvr.HandleFunc("GET /form/", getForm)
	srvr.HandleFunc("GET /dogs/", getDogs)
	srvr.HandleFunc("POST /dog", createDog)
	srvr.HandleFunc("DELETE /dog/{id}", deleteDog)
	srvr.HandleFunc("PUT /select/{id}", selectDog)

	public_dir := http.Dir("public")
	fs := http.FileServer(public_dir)
	srvr.Handle("/", http.StripPrefix("", fs))

	log.Fatal(http.ListenAndServe(":8080", srvr))
}
