package contact

import (
	"errors"

	"gorm.io/gorm"
	"hrm-app/internal/pkg/database"
)

type Repository interface {
	Create(contact *Contact) error
	FindByEmail(email string) (*Contact, error)
	FindByID(id uint) (*Contact, error)
	Update(contact *Contact) (*Contact, error)
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) Create(contact *Contact) error {
	return database.DB.Create(contact).Error
}

func (r *repository) FindByEmail(email string) (*Contact, error) {
	var contact Contact
	err := database.DB.Where("email = ?", email).First(&contact).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Jangan return error, biar handler bisa bedain antara "tidak ada data" dan "DB error"
			return &Contact{}, nil
		}
		return nil, err
	}

	return &contact, nil
}

func (r *repository) FindByID(id uint) (*Contact, error) {
	var contact Contact
	err := database.DB.First(&contact, id).Error

	return &contact, err
}

func (r *repository) Update(contact *Contact) (*Contact, error) {
	err := database.DB.Save(contact).Error
	return contact, err
}