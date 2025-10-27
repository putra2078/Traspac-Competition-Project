package contact

import (
	"errors"
)

type UseCase interface {
	Register(contact *Contact) error
	GetAll() ([]Contact, error)
	GetByID(id uint) (*Contact, error)
	DeleteByID(id uint) error
}

type usecase struct {
	repo Repository
}

func NewUseCase(repo Repository) UseCase {
	return &usecase{repo: repo}
}

func (u *usecase) Register(contact *Contact) error {
	existing, _ := u.repo.FindByEmail(contact.Email)
	if existing != nil && existing.ID != 0 {
		return errors.New("Email already in use")
	}

	return u.repo.Create(contact)
}

func (u *usecase) GetAll() ([]Contact, error) {
	return u.repo.FindAll()
}

func (u *usecase) GetByID(id uint) (*Contact, error) {
	return u.repo.FindByID(id)
}

func (u *usecase) DeleteByID(id uint) error {
	return u.repo.Delete(id)
}