package user

import (
	"hrm-app/internal/pkg/database"
	"errors"
	"gorm.io/gorm"
)

type Repository interface {
	Create(user *User) error
	FindAll() ([]User, error)
	FindByID(id uint) (*User, error)
	FindByEmail(email string) (*User, error)
	Update(user *User) error
	Delete(id uint) error
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) Create(user *User) error {
	return database.DB.Create(user).Error
}

func (r *repository) FindAll() ([]User, error) {
	var users []User
	err := database.DB.Find(&users).Error
	return users, err
}

func (r *repository) FindByID(id uint) (*User, error) {
	var user User
	err := database.DB.First(&user, id).Error

	return &user, err
}

func (r *repository) FindByEmail(email string) (*User, error) {
	var user User
	err := database.DB.Where("email = ?", email).First(&user).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Jangan return error, biar handler bisa bedain antara "tidak ada data" dan "DB error"
			return &User{}, nil
		}
		return nil, err
	}

	return &user, nil
}


func (r *repository) Update(user *User) error {
	return database.DB.Save(user).Error
}

func (r *repository) Delete(id uint) error {
	return database.DB.Delete(&User{}, id).Error
}