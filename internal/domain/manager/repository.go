package manager

import (
	"hrm-app/internal/pkg/database"
	"errors"
	"gorm.io/gorm"
)

type Repository interface {
	Create(mng *Manager) error
	FindAll() ([]Manager, error)
	FindByID(id uint) (*Manager, error)
	FindByNIP(nip string) (*Manager, error)
	FindByEmail(email string) (*Manager, error)
	Update(emp *Manager) error
	Delete(id uint) error
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) Create(emp *Manager) error {
	return database.DB.Create(emp).Error
}

func (r *repository) FindAll() ([]Manager, error) {
	var managers []Manager
	err := database.DB.Find(&managers).Error
	return managers, err
}

func (r *repository) FindByID(id uint) (*Manager, error) {
	var managers Manager
	err := database.DB.First(&managers, id).Error

	return &managers, err
}

func (r *repository) FindByNIP(nip string) (*Manager, error) {
	var managers Manager
	err := database.DB.Where("nip = ?", nip).First(&managers).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &Manager{}, nil
		}
		return nil, err
	}
	return &managers, nil
}

func (r *repository) FindByEmail(email string) (*Manager, error) {
	var managers Manager
	err := database.DB.Where("email = ?", email).First(&managers).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Jangan return error, biar handler bisa bedain antara "tidak ada data" dan "DB error"
			return &Manager{}, nil
		}
		return nil, err
	}

	return &managers, nil
}


func (r *repository) Update(Manager *Manager) error {
	return database.DB.Save(Manager).Error
}

func (r *repository) Delete(id uint) error {
	return database.DB.Delete(&Manager{}, id).Error
}