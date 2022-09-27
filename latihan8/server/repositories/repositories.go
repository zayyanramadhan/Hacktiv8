package repositories

import "sesi8-latihan/server/models"

type PersonRepo interface {
	GetPeople() (*[]models.Person, error)
	GetPeopleByID(id int) (*[]models.Person, error)
	CreatePeople(name string, address string) (id int, err error)
	UpdatePeople(name string, address string) (id int, err error)
	DeletePeople(id int) error
}
