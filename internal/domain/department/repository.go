package department

import (
	"errors"

	"hrm-app/internal/pkg/database"

	"gorm.io/gorm"
)

type Repository interface {
	Create(department *Department) error
	FindAll() ([]Department, error)
	FindByID(id uint) (*Department, error)
	FindBySlug(slug string) (*Department, error)
	Update(department *Department) error
	Delete(id uint) error
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) Create(department *Department) error {
	return database.DB.Create(department).Error
}

func (r *repository) FindAll() ([]Department, error) {
	var departments []Department
	err := database.DB.Find(&departments).Error
	return departments, err
}

func (r *repository) FindByID(id uint) (*Department, error) {
	var department Department
	err := database.DB.Find(&department, id).Error

	return &department, err
}

func (r *repository) FindBySlug(slug string) (*Department, error) {
	var department Department
	err := database.DB.Where("slug = ?", slug).First(&department).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &Department{}, nil
		}
		return nil, err
	}

	return &department, err
}

func (r *repository) Update(department *Department) error {
	return database.DB.Save(department).Error
}

func (r *repository) Delete(id uint) error {
	return database.DB.Delete(&Department{}, id).Error
}
