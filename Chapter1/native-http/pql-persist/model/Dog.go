package model

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/google/uuid"
)

type Dog struct {
	Id    string
	Name  string
	Breed string
}

type Dogs map[string]Dog

var dogMap = make(Dogs)

func AddDog(name, breed string, dogs Dogs) Dog {
	id := uuid.NewString()
	dog := Dog{
		Id:    id,
		Name:  name,
		Breed: breed,
	}

	return dog
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
