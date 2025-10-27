package employee

import (
	"errors"
)

type UseCase interface {
	Register(employee *Employee) error
	GetAll() ([]Employee, error)
	GetByID(id uint) (*Employee, error)
	DeleteByID(id uint) error
}

type usecase struct {
	repo Repository
}

func NewUseCase(repo Repository) UseCase {
	return &usecase{repo: repo}
}

func (u *usecase) Register(employee *Employee) error {
	existing, _ := u.repo.FindByNIP(employee.Nip)
	if existing != nil && existing.ID != 0 {
		return errors.New("NIP already in use")
	}

	return u.repo.Create(employee)
}

func (u *usecase) GetAll() ([]Employee, error) {
	return u.repo.FindAll()
}

func (u *usecase) GetByID(id uint) (*Employee, error) {
	return u.repo.FindByID(id)
}

func (u *usecase) DeleteByID(id uint) error {
	return u.repo.Delete(id)
}