package presence

import "time"

type Presence struct {
	ID         	  uint      `json:"id" gorm:"primaryKey"`
	EmployeeID 	  uint      `json:"employee_id"`
	Date       	  time.Time `json:"date" gorm:"type:date"`

	CheckInHour   string `json:"check_in_hour" gorm:"type:time"`
	LatCheckIn    float64   `json:"lat_check_in"`
	LongCheckIn   float64   `json:"long_check_in"`
	CheckInStatus string    `json:"check_in_status"`

	CheckOutHour   *string `json:"check_out_hour" gorm:"type:time"`
	LatCheckOut    *float64   `json:"lat_check_out"`
	LongCheckOut   *float64   `json:"long_check_out"`
	CheckOutStatus *string    `json:"check_out_status"`
}
