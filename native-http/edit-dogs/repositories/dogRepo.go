package repositories

import (
	"database/sql"

	"log"

	"github.com/fadykuzman/htmx-pql/model"
	persist "github.com/fadykuzman/htmx-pql/persistence"
)

type DogRepository struct {
	db *sql.DB
}

func NewDogRepository() *DogRepository {
	return &DogRepository{
		db: persist.GetDB(),
	}
}

func (r *DogRepository) GetDogs() ([]model.Dog, error) {

	dogs := make([]model.Dog, 0)
	rows, err := r.db.Query("SELECT * FROM dogs")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		d := model.Dog{}
		if err := rows.Scan(&d.Id, &d.Name, &d.Breed); err != nil {
			log.Fatal(err)
		}
		dogs = append(dogs, d)
	}

	return dogs, nil
}

func (r *DogRepository) CreateDog(name, breed string) model.Dog {
	dog := model.Dog{
		Name:  name,
		Breed: breed,
	}
	query := `
	INSERT INTO dogs (name, breed)
	VALUES ($1, $2)
	RETURNING id`

	r.db.QueryRow(query, name, breed).Scan(&dog.Id)
	return dog
}

func (r *DogRepository) DeleteDog(id string) {
	query := `
	DELETE FROM dogs WHERE id=$1
	`
	r.db.QueryRow(query, id)
}

func (r *DogRepository) GetDog(id string) model.Dog {
	dog := model.Dog{
		Id: id,
	}
	query := `
	SELECT name, breed FROM dogs WHERE id=($1::uuid)
    `
	r.db.QueryRow(query, id).Scan(&dog.Name, &dog.Breed)

	return dog
}

func (r *DogRepository) UpdateDog(dog model.Dog) {
	query := `
		UPDATE dogs 
	SET name=$1, breed=$2
	WHERE id=($3::uuid)
    `
	r.db.QueryRow(query, dog.Name, dog.Breed, dog.Id)
}
