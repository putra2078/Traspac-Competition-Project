package employee

import (
	"errors"
	// "time"

	"hrm-app/internal/domain/contact"
	"hrm-app/internal/domain/user"
	"hrm-app/internal/pkg/database"
	"hrm-app/internal/pkg/utils"

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

// RegisterWithContact creates a contact, employee, and user in a single DB transaction.
func (u *usecase) RegisterWithContact(employee *Employee, contactEmployee *contact.Contact, user *user.User) error {
	// simple validation on provided structs
	if employee == nil || contactEmployee == nil || user == nil {
		return errors.New("employee or contact is nil")
	}

	// use GORM transaction to ensure atomicity
	return database.DB.Transaction(func(tx *gorm.DB) error {
		// check NIP uniqueness within the transaction
		existingEmployee, err := u.repo.FindByNIP(employee.Nip)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
		if existingEmployee != nil && existingEmployee.ID != 0 {
			return errors.New("NIP already in use")
		}

		// check contact email uniqueness within the transaction
		existingContact := &contact.Contact{}
		if err := tx.Where("email = ?", contactEmployee.Email).First(existingContact).Error; err == nil {
			return errors.New("Email already in use")
		} else if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		// Create contact
		if err := tx.Create(contactEmployee).Error; err != nil {
			return err
		}


		// set user name on User and create user and employee
		user.Username = contactEmployee.Name

		// set user email on User and create user account
		user.Email = contactEmployee.Email

		// hash password before storing
		hashedPassword, err := utils.HashPassword(user.Password)
		if err != nil {
			return err
		}
		user.Password = hashedPassword

		// create user
		if err := tx.Create(user).Error; err != nil {
			return err
		}

		// set user_id on employee and create employee
		employee.UserID = user.ID

		// set contact id on employee and create employee
		employee.ContactID = contactEmployee.ID

		// ensure CreatedAt/UpdatedAt if zero (GORM will handle normally)
		// if employee.CreatedAt.IsZero() {
		// 	employee.CreatedAt = time.Now()
		// }
		// if contact.CreatedAt.IsZero() {
		// 	contact.CreatedAt = time.Now()
		// }

		if err := tx.Create(employee).Error; err != nil {
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
