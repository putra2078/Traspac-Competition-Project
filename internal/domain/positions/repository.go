package positions

import (
	"errors"
	"hrm-app/internal/pkg/database"

	"gorm.io/gorm"
)

type Repository interface {
	Create(positions *Positions) error
	FindAll() ([]Positions, error)
	FindByID(id uint) (*Positions, error)
	FindByName(name string) (*Positions, error)
	FindByDepartmentID(departmentID int) ([]Positions, error)
	Delete(id uint) error
	Update(positions *Positions) error
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) Create(positions *Positions) error {
	return database.DB.Create(positions).Error
}

func (r *repository) FindAll() ([]Positions, error) {
	var positions []Positions
	err := database.DB.Find(&positions).Error
	return positions, err
}

func (r *repository) FindByID(id uint) (*Positions, error) {
	var position Positions
	err := database.DB.First(&position, id).Error

	return &position, err
}

func (r *repository) FindByName(name string) (*Positions, error) {
	var position Positions
	err := database.DB.Where("name = ?", name).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &Positions{}, nil
		}
		return nil, err
	}

	return &position, err
}

func (r *repository) FindByDepartmentID(departmentID int) ([]Positions, error) {
	var positions []Positions
	err := database.DB.Where("department_id = ?", departmentID).Find(&positions).Error

	return positions, err
}

func (r *repository) Delete(id uint) error {
	return database.DB.Delete(&Positions{}, id).Error
}

func (r *repository) Update(positions *Positions) error {
	return database.DB.Save(positions).Error
}