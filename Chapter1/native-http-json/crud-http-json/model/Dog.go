package model

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"slices"
	"strings"

	"github.com/google/uuid"
)

type Dog struct {
	Name  string `json:"name"`
	Breed string `json:"breed"`
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
	type DogDTO struct {
		Id    string
		Name  string
		Breed string
	}
	dogsSlice := make([]DogDTO, 0, len(dogMap))
	for k, v := range dogMap {
		dog := DogDTO{
			Id:    k,
			Name:  v.Name,
			Breed: v.Breed,
		}
		dogsSlice = append(dogsSlice, dog)
	}

	slices.SortStableFunc(dogsSlice, func(a, b DogDTO) int {
		return strings.Compare(strings.ToLower(a.Name), strings.ToLower(b.Name))
	})
	tmpl := template.Must(template.ParseFiles("public/dogs.html"))
	tmpl.Execute(w, dogsSlice)
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
