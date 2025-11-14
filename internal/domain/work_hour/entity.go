package work_hour 	

import "time"

type WorkHour struct {
	ID 			uint 		`json:"id" gorm:"primaryKey"`
	Name 		string 		`json:"name"`
	StartTime	string 		`json:"start_time"`
	EndTime		string 		`json:"end_time"`
	CreatedAt 	time.Time 	`json:"created_at"`
	UpdatedAt 	time.Time 	`json:"updated_at"`
}