package admin

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

type AdminRequest struct {
	Admin struct{} `json:"admin" binding:"required"`

	Contact struct {
		Name        string `json:"name"`
		Photo       string `json:"photo"`
		Email       string `json:"email" binding:"required,email"`
		PhoneNumber string `json:"phone_number"`
		Gender      string `json:"gender"`
		Address     string `json:"address"`
		BirthDate   string `json:"birth_date"` // expected format: 2006-01-02
	} `json:"contactAdmin" binding:"required"`

	User struct {
		Name     string `json:"name"`
		Email    string `json:"email" gorm:"uniqueIndex"`
		Password string `json:"password"`
		Role     string `json:"role"`
	} `json:"user" binding:"required"`
}

func (h *Handler) RegisterWithContact(c *gin.Context) {
	var req AdminRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var birth time.Time
	if req.Contact.BirthDate != "" {
		if t, err := time.Parse("2006-01-02", req.Contact.BirthDate); err == nil {
			birth = t
		}
	}

	contactAdmin := &contact.Contact{
		Name:        req.Contact.Name,
		Photo:       req.Contact.Photo,
		Email:       req.Contact.Email,
		PhoneNumber: req.Contact.BirthDate,
		Gender:      req.Contact.Gender,
		Address:     req.Contact.Address,
		BirthDate:   birth,
	}

	admin := &Admin{}

	userAdmin := &user.User{
		Password: req.User.Password,
		Role:     req.User.Role,
	}

	if err := h.usecase.RegisterWithContact(admin, contactAdmin, userAdmin); err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Admin successfully registered"})
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

	response.Success(c, data)
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
