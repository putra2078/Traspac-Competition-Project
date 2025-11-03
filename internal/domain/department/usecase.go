package department

import (
	"errors"
)

type UseCase interface {
	Register(department *Department) error
	GetAll() ([]Department, error)
	GetByID(id uint) (*Department, error)
	GetBySlug(slug string) (*Department, error)
	DeleteByID(id uint) error
}

type usecase struct {
	repo Repository
}

func NewUseCase(repo Repository) UseCase {
	return &usecase{repo: repo}
}

func (u *usecase) Register(department *Department) error {
	existing, _ := u.repo.FindBySlug(department.Slug)
	if existing != nil && existing.ID != 0 {
		return errors.New("Slug already in use")
	}

	return u.repo.Create(department)
}

func (u *usecase) GetAll() ([]Department, error) {
	return u.repo.FindAll()
}

func (u *usecase) GetByID(id uint) (*Department, error) {
	return u.repo.FindByID(id)
}

func (u *usecase) GetBySlug(slug string) (*Department, error) {
	return u.repo.FindBySlug(slug)
}

func (u *usecase) DeleteByID(id uint) error {
	return u.repo.Delete(id)
}
