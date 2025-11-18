package positions

import (
	"errors"
)

type UseCase interface {
	Create(positions *Positions) error
	GetAll() ([]Positions, error)
	GetByID(id uint) (*Positions, error)
	GetByDepartmentID(departmentID int) ([]Positions, error)
	// Update(positions *Positions) error
	DeleteByID(id uint) error
}

type usecase struct {
	repo Repository
}

func NewUseCase(repo Repository) UseCase {
	return &usecase{repo: repo}
}

func (u *usecase) Create(positions *Positions) error {
	existing, _ := u.repo.FindByName(positions.Name)
	if existing != nil && existing.ID != 0 {
		return errors.New("Name already in use")
	}

	return u.repo.Create(positions)
}

func (u *usecase) GetAll() ([]Positions, error) {
	return u.repo.FindAll()
}

func (u *usecase) GetByID(id uint) (*Positions, error) {
	return u.repo.FindByID(id)
}

func (u *usecase) GetByDepartmentID(departmentID int) ([]Positions, error) {
	return u.repo.FindByDepartmentID(departmentID)
}

// func (u *usecase) Update(positions Positions) error {
// 	return u.repo.Update(&positions)
// }

func (u *usecase) DeleteByID(id uint) error {
	return u.repo.Delete(id)
}
