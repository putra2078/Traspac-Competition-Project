package manager

import (
	"errors"

	"hrm-app/internal/domain/contact"
	"hrm-app/internal/domain/user"
	"hrm-app/internal/pkg/database"
	"hrm-app/internal/pkg/utils"

	"gorm.io/gorm"
)

type UseCase interface {
	Register(manager *Manager) error
	RegisterWithContact(manager *Manager, contact *contact.Contact, user *user.User) error
	GetAll() ([]Manager, error)
	GetByID(id uint) (*Manager, error)
	DeleteByID(id uint) error
}

type usecase struct {
	repo Repository
}

func NewUseCase(repo Repository) UseCase {
	return &usecase{repo: repo}
}

func (u *usecase) Register(managers *Manager) error {
	existing, _ := u.repo.FindByNIP(managers.Nip)
	if existing != nil && existing.ID != 0 {
		return errors.New("NIP already in use")
	}

	return u.repo.Create(managers)
}

// RegisterWithContact creates a contact and manager in a single DB transaction.
// This implementation uses GORM transaction directly so we don't need to
// change repository signatures. It guarantees atomicity: either both rows
// are inserted or none.
func (u *usecase) RegisterWithContact(manager *Manager, cont *contact.Contact, usr *user.User) error {
	if manager == nil || cont == nil || usr == nil {
		return errors.New("manager, contact, or user is nil")
	}

	return database.DB.Transaction(func(tx *gorm.DB) error {
		// Check NIP uniqueness using transaction-scoped repo if possible
		existingManager := &Manager{}
		if err := tx.Where("nip = ?", manager.Nip).First(existingManager).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
		if existingManager.ID != 0 {
			return errors.New("NIP already in use")
		}

		// Check contact email uniqueness
		existingContact := &contact.Contact{}
		if err := tx.Where("email = ?", cont.Email).First(existingContact).Error; err == nil {
			return errors.New("contact email already in use")
		} else if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		// Create contact
		if err := tx.Create(cont).Error; err != nil {
			return err
		}

		// Set user fields and hash password
		usr.Username = cont.Name
		usr.Email = cont.Email
		hashed, err := utils.HashPassword(usr.Password)
		if err != nil {
			return err
		}
		usr.Password = hashed

		// Create user
		if err := tx.Create(usr).Error; err != nil {
			return err
		}

		// Set foreign keys and create manager
		manager.UserID = usr.ID
		manager.ContactID = cont.ID
		if err := tx.Create(manager).Error; err != nil {
			return err
		}

		return nil
	})
}

func (u *usecase) GetAll() ([]Manager, error) {
	return u.repo.FindAll()
}

func (u *usecase) GetByID(id uint) (*Manager, error) {
	return u.repo.FindByID(id)
}

func (u *usecase) DeleteByID(id uint) error {
	return u.repo.Delete(id)
}
