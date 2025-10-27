package user

import (
	"errors"
	"hrm-app/internal/pkg/utils"
)

type UseCase interface {
	Register(user *User) error
	GetAll() ([]User, error)
	GetByID(id uint) (*User, error)
	DeleteByID(id uint) error
}

type usecase struct {
	repo Repository
}

func NewUseCase(repo Repository) UseCase {
	return &usecase{repo: repo}
}

func (u *usecase) Register(user *User) error {
	existing, _ := u.repo.FindByEmail(user.Email)
	if existing != nil && existing.ID != 0 {
		return errors.New("Email already in use")
	}

	hashed, _ := utils.HashPassword(user.Password)
	user.Password = hashed
	return u.repo.Create(user)
}

func (u *usecase) GetAll() ([]User, error) {
	return u.repo.FindAll()
}

func (u *usecase) GetByID(id uint) (*User, error) {
	return u.repo.FindByID(id)
}

func (u *usecase) DeleteByID(id uint) error {
	return u.repo.Delete(id)
}