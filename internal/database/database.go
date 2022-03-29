package database

type CRUD interface {
	All() (interface{}, error)
	Find(id int) (interface{}, error)
	Create(u interface{}) (interface{}, error)
	Update(id int, u interface{}) (interface{}, error)
	Delete(id int) error
}
