package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/google/uuid"
)

type Dog struct {
	Id    string
	Name  string
	Breed string
}

type Dogs map[string]Dog

var dogMap = Dogs{}

func dogs(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("public/dogs.html"))
	tmpl.Execute(w, dogMap)
}

func AddDog(name, breed string, dogs Dogs) {
	dog := Dog{
		Id:    uuid.NewString(),
		Name:  name,
		Breed: breed,
	}
	dogs[dog.Id] = dog
}

func main() {
	AddDog("Rocky", "Whippet", dogMap)
	AddDog("Uranu", "Chesterfield", dogMap)

	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", http.StripPrefix("", fs))
	http.HandleFunc("/dogs", dogs)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
