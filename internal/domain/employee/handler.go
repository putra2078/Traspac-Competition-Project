package employee

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"hrm-app/internal/domain/contact"
	"hrm-app/internal/domain/user"
	"hrm-app/internal/response"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Handler struct {
	usecase UseCase
}

func NewHandler(u UseCase) *Handler {
	return &Handler{usecase: u}
}

// DTO for combined request
type EmployeeWithContactRequest struct {
	Employee struct {
		Nip          string `json:"nip" binding:"required"`
		Status       string `json:"status"`
		ManagerID    uint   `json:"manager_id"`
		PositionID   uint   `json:"position_id"`
		DepartmentID uint   `json:"department_id"`
		WorkTime     uint   `json:"work_time"`
	} `json:"employee" binding:"required"`

	Contact struct {
		Name        string `json:"name" binding:"required"`
		Photo       string `json:"photo"`
		Email       string `json:"email" binding:"required,email"`
		PhoneNumber string `json:"phone_number"`
		Gender      string `json:"gender"`
		Address     string `json:"address"`
		BirthDate   string `json:"birth_date"` // expected format: 2006-01-02
	} `json:"contact" binding:"required"`

	User struct {
		Name     string `json:"name"`
		Email    string `json:"email" gorm:"uniqueIndex"`
		Password string `json:"password"`
		Role     string `json:"role"`
	} `json:"user" binding:"required"`
}

func (h *Handler) Register(c *gin.Context) {
	var employee Employee
	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.usecase.Register(&employee); err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Employee registered successfully"})
}

// RegisterWithContact handles a single JSON body that contains both employee
// and contact data. It maps the DTO to internal entities and delegates to
// usecase.RegisterWithContact which runs a DB transaction.
func (h *Handler) RegisterWithContact(c *gin.Context) {
	var req EmployeeWithContactRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// map to contact.Contact
	var birth time.Time
	if req.Contact.BirthDate != "" {
		if t, err := time.Parse("2006-01-02", req.Contact.BirthDate); err == nil {
			birth = t
		}
	}

	contactEmployee := &contact.Contact{
		Name:        req.Contact.Name,
		Photo:       req.Contact.Photo,
		Email:       req.Contact.Email,
		PhoneNumber: req.Contact.PhoneNumber,
		Gender:      req.Contact.Gender,
		Address:     req.Contact.Address,
		BirthDate:   birth,
	}

	// map to employee.Employee
	employee := &Employee{
		Nip:          req.Employee.Nip,
		Status:       req.Employee.Status,
		ManagerID:    req.Employee.ManagerID,
		PositionID:   req.Employee.PositionID,
		DepartmentID: req.Employee.DepartmentID,
		WorkTime:     req.Employee.WorkTime,
	}

	user := &user.User{
		Password: req.User.Password,
		Role:     req.User.Role,
	}

	if err := h.usecase.RegisterWithContact(employee, contactEmployee, user); err != nil {
		// choose response code depending on error - for simplicity return 409 for conflicts
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "employee and contact created"})
}

func (h *Handler) GetAll(c *gin.Context) {
	data, err := h.usecase.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, data)
}

func (h *Handler) GetByID(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil || id < 0 {
		response.Error(c, http.StatusBadRequest, "Invalid id parameter")
		return
	}

	data, err := h.usecase.GetByID(uint(id))
	if err != nil {
		response.Error(c, http.StatusNotFound, "User not found")
		return
	}

	response.GetSuccess(c, data)
}

func (h *Handler) Delete(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	// Check is the id parameter is not a negative int
	if err != nil || id < 0 {
		response.Error(c, http.StatusBadRequest, "Invalid id parameter")
		return
	}

	err = h.usecase.DeleteByID(uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.Error(c, http.StatusNotFound, "Data not found")
			return
		}
		response.Error(c, http.StatusInternalServerError, "failed to delete record")
		return
	}

	response.DeleteSuccess(c, "Manager deleted successfully")
}
