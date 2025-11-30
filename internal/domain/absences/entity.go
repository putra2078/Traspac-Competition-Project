package absences

import "time"

type Absences struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	EmployeeID   uint      `json:"employee_id"`
	CategoryID   uint      `json:"category_id"`
	Date         time.Time `json:"date" gorm:"type:date"`
	Description  string    `json:"description"`
	Status       string    `json:"status"`
	RejectReason string    `json:"reject_reason"`
}
