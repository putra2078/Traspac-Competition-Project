package work_hour

import (
	"errors"
)

type UseCase interface {
	Register(workHour *WorkHour) error
	GetAll() ([]WorkHour, error)
	GetByID(id uint) (*WorkHour, error)
	DeleteByID(id uint) error
}

type usecase struct {
	repo Repository
}


func NewUseCase(repo Repository) UseCase {
	return &usecase{repo: repo}
}


func (u *usecase) Register(workHour *WorkHour) error {
	// Function pengecekan duplikasi nama
	existing, _ := u.repo.FindByName(workHour.Name)
	if existing != nil && existing.ID != 0 {
		return errors.New("Name already in use")
	}
	return u.repo.Create(workHour)
}


func (u *usecase) GetAll() ([]WorkHour, error) {
	return u.repo.FindAll()
}


func (u *usecase) GetByID(id uint) (*WorkHour, error) {
	return u.repo.FindByID(id)
}


func (u *usecase) DeleteByID(id uint) error {
	return u.repo.Delete(id)
}
