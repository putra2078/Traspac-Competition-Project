package employee

import (
	"errors"

	"gorm.io/gorm"
	"hrm-app/internal/pkg/database"
)

type Repository interface {
	Create(employee *Employee) error
	FindAll() ([]Employee, error)
	FindByID(id uint) (*Employee, error)
	FindByUserID(userID uint) (*Employee, error)
	FindByNIP(nip string) (*Employee, error)
	FindByEmail(email string) (*Employee, error)
	Update(emp *Employee) error
	Delete(id uint) error
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) Create(employee *Employee) error {
	return database.DB.Create(employee).Error
}

func (r *repository) FindAll() ([]Employee, error) {
	var employees []Employee
	err := database.DB.Find(&employees).Error
	return employees, err
}

func (r *repository) FindByID(id uint) (*Employee, error) {
	var employee Employee
	err := database.DB.First(&employee, id).Error

	return &employee, err
}

func (r *repository) FindByNIP(nip string) (*Employee, error) {
	var employee Employee
	err := database.DB.Where("nip = ?", nip).First(&employee).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &Employee{}, nil
		}
		return nil, err
	}
	return &employee, nil
}

func (r *repository) FindByEmail(email string) (*Employee, error) {
	var employee Employee
	err := database.DB.Where("email = ?", email).First(&employee).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Jangan return error, biar handler bisa bedain antara "tidak ada data" dan "DB error"
			return &Employee{}, nil
		}
		return nil, err
	}

	return &employee, nil
}

func (r *repository) FindByUserID(userID uint) (*Employee, error) {
	var employee Employee
	err := database.DB.Where("user_id = ?", userID).First(&employee).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &Employee{}, nil
		}
		return nil, err
	}
	return &employee, nil
}

func (r *repository) Update(employee *Employee) error {
	return database.DB.Save(employee).Error
}

func (r *repository) Delete(id uint) error {
	return database.DB.Delete(&Employee{}, id).Error
}
