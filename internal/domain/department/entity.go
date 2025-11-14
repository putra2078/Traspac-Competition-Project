package department

import "time"

type Department struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name"`
	Slug        string    `json:"slug"`
	ManagerID   *uint      `json:"manager_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
