package employee

import (
	"hrm-app/internal/domain/contact"
	"hrm-app/internal/domain/user"
	"time"
)

type Employee struct {
	ID			    uint      `json:"id" gorm:"primaryKey"`
	UserID  	    uint      `json:"user_id" gorm:"uniqueIndex"`
	User 			user.User `gorm:"constraint:OnUpdate:CASCADE.OnDelete:CASCADE;"`
	Nip	    	    string    `json:"nip" gorm:"uniqueIndex"`
	Status		    string    `json:"status"`
	ManagerID		uint      `json:"manager_id"`
	ContactID       uint      `json:"contact_id"`
	Contact         contact.Contact `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	PositionID	    uint      `json:"position_id"`
	DepartmentID 	uint      `json:"department_id"`
	CreatedAt		time.Time `json:"created_at"`
	UpdatedAt		time.Time `json:"updated_at"`
}
