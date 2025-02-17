package model

import (
	"github.com/google/uuid"
)

type Dog struct {
	Id    string
	Name  string
	Breed string
}

type DogMap map[string]Dog

func addDog(dogMap DogMap, name, breed string) Dog {
	id := uuid.NewString()
	dog := Dog{
		Id:    id,
		Name:  name,
		Breed: breed,
	}
	dogMap[id] = dog
	return dog
}
