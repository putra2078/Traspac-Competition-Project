package work_hour

import (
	// "errors"

	// "gorm.io/gorm"
	"hrm-app/internal/pkg/database"
)

type Repository interface {
	Create(workHour *WorkHour) error
	FindAll() ([]WorkHour, error)
	FindByID(id uint) (*WorkHour, error)
	FindByName(name string) (*WorkHour, error)
	Update(workHour *WorkHour) error
	Delete(id uint) error
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) Create(workHour *WorkHour) error {
	return database.DB.Create(workHour).Error
}

func (r *repository) FindAll() ([]WorkHour, error) {
	var workHour []WorkHour
	err := database.DB.Find(&workHour).Error
	return workHour, err
}

func (r *repository) FindByID(id uint) (*WorkHour, error) {
	var workHour WorkHour
	err := database.DB.First(&workHour, id).Error

	return &workHour, err
}

func (r *repository) FindByName(name string) (*WorkHour, error) {
	var workHour WorkHour
	err := database.DB.Where("name = ?", name).First(&workHour).Error

	return &workHour, err
}

func (r *repository) Update(workHour *WorkHour) error {
	return database.DB.Save(workHour).Error
}

func (r *repository) Delete(id uint) error {
	return database.DB.Delete(&WorkHour{}, id).Error
}
