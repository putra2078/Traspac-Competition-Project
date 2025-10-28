package employee

import (
	"errors"
	"time"

	"hrm-app/internal/domain/contact"
	"hrm-app/internal/domain/user"
	"hrm-app/internal/pkg/database"

	"gorm.io/gorm"
)

type UseCase interface {
	Register(employee *Employee) error
	RegisterWithContact(employee *Employee, contact *contact.Contact, user *user.User) error
	GetAll() ([]Employee, error)
	GetByID(id uint) (*Employee, error)
	DeleteByID(id uint) error
}

type usecase struct {
	repo Repository
}

func NewUseCase(repo Repository) UseCase {
	return &usecase{repo: repo}
}

func (u *usecase) Register(employee *Employee) error {
	existing, _ := u.repo.FindByNIP(employee.Nip)
	if existing != nil && existing.ID != 0 {
		return errors.New("NIP already in use")
	}

	return u.repo.Create(employee)
}

// RegisterWithContact creates a contact and employee in a single DB transaction.
// This implementation uses GORM transaction directly so we don't need to
// change repository signatures. It guarantees atomicity: either both rows
// are inserted or none.
func (u *usecase) RegisterWithContact(emp *Employee, cont *contact.Contact, usr *user.User) error {
	// simple validation on provided structs
	if emp == nil || cont == nil || usr == nil {
		return errors.New("employee or contact is nil")
	}

	// use GORM transaction to ensure atomicity
	return database.DB.Transaction(func(tx *gorm.DB) error {
		// check NIP uniqueness within the transaction
		var existingEmp Employee
		if err := tx.Where("nip = ?", emp.Nip).First(&existingEmp).Error; err == nil {
			return errors.New("NIP already in use")
		} else if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		// check contact email uniqueness within the transaction
		var existingContact contact.Contact
		if err := tx.Where("email = ?", cont.Email).First(&existingContact).Error; err == nil {
			return errors.New("contact email already in use")
		} else if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		// create contact
		if err := tx.Create(cont).Error; err != nil {
			return err
		}

		// set user name on User and create user and employee
		usr.Name = cont.Name

		// set user email on User and create user account
		usr.Email = cont.Email

		// create user
		if err := tx.Create(usr).Error; err != nil {
			return err
		}

		// set user_id on employee and create employee
		emp.UserID = usr.ID

		// set contact id on employee and create employee
		emp.ContactID = cont.ID

		// ensure CreatedAt/UpdatedAt if zero (GORM will handle normally)
		if emp.CreatedAt.IsZero() {
			emp.CreatedAt = time.Now()
		}
		if cont.CreatedAt.IsZero() {
			cont.CreatedAt = time.Now()
		}

		if err := tx.Create(emp).Error; err != nil {
			return err
		}

		return nil
	})
}

func (u *usecase) GetAll() ([]Employee, error) {
	return u.repo.FindAll()
}

func (u *usecase) GetByID(id uint) (*Employee, error) {
	return u.repo.FindByID(id)
}

func (u *usecase) DeleteByID(id uint) error {
	return u.repo.Delete(id)
}
