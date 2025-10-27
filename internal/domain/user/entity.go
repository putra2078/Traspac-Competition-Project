package user

import "time"

type User struct {
	ID		uint      `json:"id" gorm:"primaryKey"`
	Name	string    `json:"name"`
	Email   string    `json:"email" gorm:"uniqueIndex"`
	Password string   `json:"password"`
	Role    string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}