package model

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/google/uuid"
)

type Dog struct {
	Name  string
	Breed string
}

type Dogs map[string]Dog

var dogMap = make(Dogs)

func AddDog(name, breed string, dogs Dogs) Dog {
	id := uuid.NewString()
	dog := Dog{
		Name:  name,
		Breed: breed,
	}
	dogs[id] = dog
	data, err := json.MarshalIndent(dogs, "", " ")
	if err != nil {
		log.Println(err)
	}

	os.WriteFile("resources/dogs.json", data, 0644)

	return dog
}

func GetDogs(w http.ResponseWriter, r *http.Request) {
}

func CreateDog(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	breed := r.FormValue("breed")
	dog := AddDog(name, breed, dogMap)
	slice := []Dog{
		dog,
	}
	tmpl := template.Must(template.ParseFiles("public/dogs.html"))
	tmpl.Execute(w, slice)

}

func DeleteDog(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	fmt.Println(id)
	delete(dogMap, id)
}

func ReadDogsFromFile(path string) {
	file, err := os.ReadFile(path)
	if err != nil {
		log.Println(err)
	}
	var rawDogs map[string]Dog

	if err := json.Unmarshal(file, &rawDogs); err != nil {
		log.Println(err)
	}

	for k, v := range rawDogs {
		dogMap[k] = v
	}
}
