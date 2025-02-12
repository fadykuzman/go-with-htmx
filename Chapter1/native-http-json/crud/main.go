package main

import (
	"log"
	"net/http"

	"github.com/fadykuzman/htmx-http-json/model"
)

func main() {

	model.ReadDogsFromFile("resources/dogs.json")

	r := http.NewServeMux()
	r.HandleFunc("GET /dogs/", model.GetDogs)
	r.HandleFunc("POST /dog", model.CreateDog)
	r.HandleFunc("DELETE /dog/{id}", model.DeleteDog)

	fs := http.FileServer(http.Dir("public"))
	r.Handle("/", http.StripPrefix("", fs))

	log.Fatal(http.ListenAndServe(":8081", r))
}
