package presence

import (
	// "errors"
	"time"

	"hrm-app/internal/pkg/database"
	// "gorm.io/gorm"
)

type Repository interface {
	CheckIn(checkIn *Presence) error
	CheckOut(presence *Presence) error
	FindAll() ([]Presence, error)
	FindByEmployeeDateRange(employeeID uint, startDate, endDate time.Time) ([]Presence, error)
	FilterByDateRange(startDate, endDate time.Time) ([]Presence, error)
	// FindCheckinToday(employeeID uint) (*Presence, error)
	FindCheckinToday(employeeID uint) (*Presence, error)
	FindCheckOutToday(employeeID uint) (*Presence, error)
	CreateCheckin(presence *Presence) error
	Delete(id uint) error
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) CheckIn(checkIn *Presence) error {
	return database.DB.Create(checkIn).Error
}

func (r *repository) CheckOut(p *Presence) error {
	return database.DB.Model(&Presence{}).
		Where("id = ?", p.ID).
		Updates(map[string]interface{}{
			"check_out_hour":   p.CheckOutHour,
			"lat_check_out":    p.LatCheckOut,
			"long_check_out":   p.LongCheckOut,
			"check_out_status": p.CheckOutStatus,
		}).Error
}

func (r *repository) FindAll() ([]Presence, error) {
	var presences []Presence
	err := database.DB.Find(&presences).Error
	return presences, err
}

func (r *repository) FindByEmployeeID(employeeID uint) ([]Presence, error) {
	var presences []Presence
	if err := database.DB.Where("employee_id = ?", employeeID).Find(&presences).Error; err != nil {
		return nil, err
	}

	return presences, nil
}

func (r *repository) FilterByDateRange(startDate, endDate time.Time) ([]Presence, error) {
	var presences []Presence
	if err := database.DB.Where("date BETWEEN ? and ?", startDate, endDate).Find(&presences).Error; err != nil {
		return nil, err
	}

	return presences, nil
}

func (r *repository) FindByEmployeeDateRange(employeeID uint, startDate, endDate time.Time) ([]Presence, error) {
	var presences []Presence
	if err := database.DB.Where("employee_id = ?", employeeID).
		Where("date BETWEEN ? and ?", startDate, endDate).Find(&presences).Error; err != nil {
		return nil, err
	}

	return presences, nil
}

func (r *repository) Delete(id uint) error {
	return database.DB.Delete(&Presence{}, id).Error
}

func (r *repository) FindCheckinToday(employeeID uint) (*Presence, error) {
	var presence Presence
	today := time.Now().Format("2006-01-02")

	err := database.DB.
		Where("employee_id = ? AND DATE(date) = ?", employeeID, today).
		First(&presence).Error
	if err != nil {
		return nil, err
	}
	return &presence, nil
}

func (r *repository) FindCheckOutToday(employeeID uint) (*Presence, error) {
	var presence Presence

	err := database.DB.
		Where("employee_id = ?", employeeID).
		Where("date = CURRENT_DATE").
		Where("(check_out_status = '' OR check_out_status IS NULL)").
		First(&presence).Error
	if err != nil {
		return nil, err
	}

	return &presence, nil
}

func (r *repository) CreateCheckin(presence *Presence) error {
	return database.DB.Create(presence).Error
}
