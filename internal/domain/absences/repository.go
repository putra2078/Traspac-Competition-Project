package absences 	

import (
	"hrm-app/internal/pkg/database"
)

type Repository interface {
	Create(absence *Absences) error
	FindAll() ([]Absences, error)
	FindByID(id uint) (*Absences, error)
	Update(absence *Absences) error
	Delete(id uint) error
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) Create(absence *Absences) error {
	return database.DB.Create(absence).Error
}

func (r *repository) FindAll() ([]Absences, error) {
	var absences []Absences
	err := database.DB.Find(&absences).Error

	return absences, err
}

func (r *repository) FindByID(id uint) (*Absences, error) {
	var absence Absences
	err := database.DB.First(&absence, id).Error

	return &absence, err
}

func (r *repository) FindByEmployeeID(employeeID uint) (*Absences, error) {
	var absence Absences
	err := database.DB.Where("employee_id = ?", employeeID).Error

	return &absence, err
} 

func (r *repository) Update(absence *Absences) error {
	return database.DB.Save(absence).Error
}

func (r *repository) Delete(id uint) error {
	return database.DB.Delete(&Absences{}, id).Error
}