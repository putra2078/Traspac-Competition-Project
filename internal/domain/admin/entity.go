package admin

import "time"

type Admin struct {
	ID 			uint 		`json:"id" gorm:"primaryKey"`
	UserID 		uint 		`json:"user_id"`
	ContactID 	uint 		`json:"contact_id"`
	CreatedAt 	time.Time 	`json:"created_at"`
	UpdatedAt 	time.Time 	`json:"updated_at"`
}