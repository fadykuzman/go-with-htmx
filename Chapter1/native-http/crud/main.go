package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type Dog struct {
	Id    string
	Name  string
	Breed string
}

type Dogs map[string]Dog

var dogMap = make(Dogs)

func GetDogs(w http.ResponseWriter, r *http.Request) {

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

func createDog(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	breed := r.FormValue("breed")
	AddDog(name, breed, dogMap)
}

func deleteDog(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Println(id)
	delete(dogMap, id)
}

func main() {
	r := mux.NewRouter().StrictSlash(true)
	AddDog("Rocky", "Whippet", dogMap)
	AddDog("Uranu", "Chesterfield", dogMap)
	fs := http.FileServer(http.Dir("public"))
	r.HandleFunc("/dogs/", GetDogs).Methods("GET")
	r.HandleFunc("/dog", createDog).Methods("POST")
	r.HandleFunc("/dog/{id}", deleteDog).Methods("DELETE")

	r.PathPrefix("/").Handler(http.StripPrefix("", fs))
	log.Fatal(http.ListenAndServe(":8081", r))
}
