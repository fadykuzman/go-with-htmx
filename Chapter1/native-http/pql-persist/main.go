package main

import (
	"log"
	"net/http"

	"github.com/fadykuzman/htmx-pql/model"
	persist "github.com/fadykuzman/htmx-pql/persistence"
)

func main() {
	srvr := http.NewServeMux()

	srvr.HandleFunc("GET /dogs/", model.GetDogs)
	srvr.HandleFunc("POST /dog", model.CreateDog)
	srvr.HandleFunc("DELETE /dog/{id}", model.DeleteDog)
	public_dir := http.Dir("public")
	fs := http.FileServer(public_dir)
	srvr.Handle("/", http.StripPrefix("", fs))

	persist.StartDB()

	log.Fatal(http.ListenAndServe(":8080", srvr))
}
