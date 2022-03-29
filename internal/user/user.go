package user

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

type User struct {
	gorm.Model
	Name string
}

func New(file string) (*Repository, error) {
	db, err := gorm.Open(sqlite.Open(file), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &Repository{db}, nil
}

func (r *Repository) All() (interface{}, error) {
	return []byte("tests"), nil
}

func (r *Repository) Find(id int) (interface{}, error) {
	panic("not implemented") // TODO: Implement
}

func (r *Repository) Create(u interface{}) (interface{}, error) {
	panic("not implemented") // TODO: Implement
}

func (r *Repository) Update(id int, u interface{}) (interface{}, error) {
	panic("not implemented") // TODO: Implement
}

func (r *Repository) Delete(id int) error {
	panic("not implemented") // TODO: Implement
}
