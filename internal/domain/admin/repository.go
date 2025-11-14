package admin

import (
	"hrm-app/internal/pkg/database"
)

type Repository interface {
	Create(admin *Admin) error
	FindAll() ([]Admin, error)
	FindByID(id uint) (*Admin, error)
	Update(adm *Admin) error
	Delete(id uint) error
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) Create(admin *Admin) error {
	return database.DB.Create(admin).Error
}

func (r *repository) FindAll() ([]Admin, error) {
	var admins []Admin
	err := database.DB.Find(&admins).Error

	return admins, err
}

func (r *repository) FindByID(id uint) (*Admin, error) {
	var admin Admin
	err := database.DB.First(&admin, id).Error

	return &admin, err
}

func (r *repository) Update(admin *Admin) error {
	return database.DB.Save(admin).Error
}

func (r *repository) Delete(id uint) error {
	return database.DB.Delete(&Admin{}, id).Error
}