package department

import "time"

type Department struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name" gorm:"uniqueIndex"`
	Slug        string    `json:"slug" gorm:"uniqueIndex"`
	HeadManager uint      `json:"head_manager"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
