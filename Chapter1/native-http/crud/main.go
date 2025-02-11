package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"slices"
	"strings"

	"github.com/google/uuid"
)

type Dog struct {
	Id    string
	Name  string
	Breed string
}

type Dogs map[string]Dog

var dogMap = make(Dogs)

func GetDogs(w http.ResponseWriter, r *http.Request) {
	dogsSlice := make([]Dog, 0, len(dogMap))
	for _, v := range dogMap {
		dogsSlice = append(dogsSlice, v)
	}
	slices.SortStableFunc(dogsSlice, func(a, b Dog) int {
		return strings.Compare(strings.ToLower(a.Name), strings.ToLower(b.Name))
	})
	tmpl := template.Must(template.ParseFiles("public/dogs.html"))
	tmpl.Execute(w, dogsSlice)
}

func AddDog(name, breed string, dogs Dogs) Dog {
	dog := Dog{
		Id:    uuid.NewString(),
		Name:  name,
		Breed: breed,
	}
	dogs[dog.Id] = dog
	return dog
}

func createDog(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	breed := r.FormValue("breed")
	dog := AddDog(name, breed, dogMap)
	slice := []Dog{
		dog,
	}
	tmpl := template.Must(template.ParseFiles("public/dogs.html"))
	tmpl.Execute(w, slice)

}

func deleteDog(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	fmt.Println(id)
	delete(dogMap, id)
}

func main() {
	AddDog("Rocky", "Whippet", dogMap)
	AddDog("Uranu", "Chesterfield", dogMap)

	r := http.NewServeMux()
	r.HandleFunc("GET /dogs/", GetDogs)
	r.HandleFunc("POST /dog", createDog)
	r.HandleFunc("DELETE /dog/{id}", deleteDog)

	fs := http.FileServer(http.Dir("public"))
	r.Handle("/", http.StripPrefix("", fs))

	log.Fatal(http.ListenAndServe(":8081", r))
}
