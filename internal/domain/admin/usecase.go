package admin

import (
	"errors"

	"hrm-app/internal/domain/contact"
	"hrm-app/internal/domain/user"
	"hrm-app/internal/pkg/database"
	"hrm-app/internal/pkg/utils"

	"gorm.io/gorm"
)

type UseCase interface {
	RegisterWithContact(admin *Admin, contactAdmin *contact.Contact, userAdmin *user.User) error
	GetAll() ([]Admin, error)
	GetByID(id uint) (*Admin, error)
	DeleteByID(id uint) error
}

type usecase struct {
	repo Repository
}

func NewUseCase(repo Repository) UseCase {
	return &usecase{repo: repo}
}

func (u *usecase) RegisterWithContact(admin *Admin, contactAdmin *contact.Contact, userAdmin *user.User) error {
	if admin == nil || contactAdmin == nil || userAdmin == nil {
		return errors.New("Admin or contact is nil")
	}

	return database.DB.Transaction(func(tx *gorm.DB) error {
		existingContact := &contact.Contact{}
		if err := tx.Where("email = ?", contactAdmin.Email).First(existingContact).Error; err == nil {
			return errors.New("Contact email already in use")
		} else if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		if err := tx.Create(contactAdmin).Error; err != nil {
			return err
		}

		userAdmin.Username = contactAdmin.Name
		userAdmin.Email = contactAdmin.Email
		hashed, err := utils.HashPassword(userAdmin.Password)
		if err != nil {
			return err
		}
		userAdmin.Password = hashed

		if err := tx.Create(userAdmin).Error; err != nil {
			return err
		}

		admin.UserID = userAdmin.ID
		admin.ContactID = contactAdmin.ID
		if err := tx.Create(admin).Error; err != nil {
			return err
		}

		return nil
	})
}

func (u *usecase) GetAll() ([]Admin, error) {
	return u.repo.FindAll()
}

func (u *usecase) GetByID(id uint) (*Admin, error) {
	return u.repo.FindByID(id)
}

func (u *usecase) DeleteByID(id uint) error {
	return u.repo.Delete(id)
}
