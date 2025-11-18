package positions 	

import "time"

type Positions struct {
	ID uint `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
	DepartmentID int `json:"department_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}