package repositories

import (
	"database/sql"
	"fmt"

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

	// dogs := make(model.Dogs)
	rows, err := r.db.Query("SELECT * FROM dogs")
	if err != nil {
		return nil, err
	}

	fmt.Println(rows)

	return make([]model.Dog, 0), nil
}
