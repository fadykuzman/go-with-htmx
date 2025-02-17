package main

import (
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

func getDogs(w http.ResponseWriter, r *http.Request) {

	dogsSlice, err := dogRepo.GetDogs()
	if err != nil {
		log.Fatal(err)
	}

	slices.SortStableFunc(dogsSlice, func(a, b model.Dog) int {
		return strings.Compare(strings.ToLower(a.Name), strings.ToLower(b.Name))
	})
	tmpl := template.Must(template.ParseFiles("public/dogs.html"))
	tmpl.Execute(w, dogsSlice)
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

func main() {
	err := persist.Connect()
	if err != nil {
		log.Println(err)
	}

	defer persist.GetDB().Close()

	dogRepo = repositories.NewDogRepository()
	srvr := http.NewServeMux()

	srvr.HandleFunc("GET /dogs/", getDogs)
	srvr.HandleFunc("POST /dog", createDog)
	srvr.HandleFunc("DELETE /dog/{id}", deleteDog)
	public_dir := http.Dir("public")
	fs := http.FileServer(public_dir)
	srvr.Handle("/", http.StripPrefix("", fs))

	log.Fatal(http.ListenAndServe(":8080", srvr))
}
