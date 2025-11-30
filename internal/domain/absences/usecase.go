package absences 	

// import (
// 	"errors"

// 	"hrm-app/internal/pkg/database"

// 	"gorm.io/gorm"
// )

// type UseCase interface {
// 	Create(absence *Absences) error
// 	GetAll() ([]Absences, error)
// 	GetByID(id uint) (*Absences, error)
// 	DeleteByID(id uint) error
// }

// type usecase struct {
// 	repo repository
// }

// func NewUseCase(repo Repository) UseCase {
// 	return &usecase{repo: repo}
// }

// func (u *usecase) Create(absence *Absences) error {
// 	existing, _ := 
// }