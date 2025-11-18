package positions 	

import (
	"errors"
	"net/http"
	"strconv"

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

func (h *Handler) Create(c *gin.Context) {
	var positions Positions
	if err :=  c.ShouldBindJSON(&positions); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.usecase.Create(&positions); err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	response.Success(c, "Position successfully created")
}

func (h *Handler) GetAll(c *gin.Context) {
	data, err := h.usecase.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response.Success(c, data)
}

func (h *Handler) GetByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil || id < 0 {
		response.Error(c, http.StatusBadRequest ,"Invalid parameter ID")
		return
	}

	data, err := h.usecase.GetByID(uint(id))
	if err != nil {
		response.Error(c, http.StatusNotFound, "Position not found")
		return
	}

	response.Success(c, data)
}

func (h *Handler) GetByDepartmentID(c *gin.Context) {
	departmentID := c.Param("departmentid")
	id, err := strconv.Atoi(departmentID)
	if err != nil || id < 0 {
		response.Error(c, http.StatusBadRequest, "Invalid parameter ID")
	}

	data, err := h.usecase.GetByDepartmentID(int(id))
	if err != nil {
		response.Error(c, http.StatusNotFound, "Position in this department ID not found")
		return
	}

	response.Success(c, data)
}

func (h *Handler) Delete(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil || id < 0 {
		response.Error(c, http.StatusBadRequest, "Invalid ID parameter")
		return
	}

	err = h.usecase.DeleteByID(uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.Error(c, http.StatusNotFound, "Data not found")
			return
		}
		response.Error(c, http.StatusInternalServerError, "failed to delete data")
		return
	}

	response.DeleteSuccess(c, "Positions successfully deleted")
}