package contact

// import (
// 	"hrm-app/internal/pkg/database"
// 	"errors"
// 	"gorm.io/gorm"
// )


type Repository interface {
	Create(contact *Contact) error
	FindAll() ([]Contact, error)
	FindByID(id uint) (*Contact, error)
	FindByEmail(email string) (*Contact, error)
	Update(contact *Contact) error
	Delete(id uint) error
}

