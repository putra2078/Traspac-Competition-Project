package employee

import "time"

type Employee struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	UserID       uint      `json:"user_id" gorm:"uniqueIndex"`
	Nip          string    `json:"nip"`
	Status       string    `json:"status"`
	ManagerID    uint      `json:"manager_id"`
	ContactID    uint      `json:"contact_id"`
	PositionID   uint      `json:"position_id"`
	DepartmentID uint      `json:"department_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
