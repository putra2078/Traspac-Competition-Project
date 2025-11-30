package work_hour

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
	var workHour WorkHour
	if err := c.ShouldBindJSON(&workHour); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err :=h.usecase.Register(&workHour); err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	response.Success(c, "Work hour successfully created")
}

func (h *Handler) GetAll(c *gin.Context) {
	data, err := h.usecase.GetAll()
	if err != nil {
		response.ErrorInternalServer(c, "Failed to get data")
		return
	}

	response.GetSuccess(c, data)
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
		response.Error(c, http.StatusNotFound, "Work hour not found")
		return
	}

	response.GetSuccess(c, data)
}

func (h *Handler) Delete(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil || id < 0 {
		response.Error(c, http.StatusBadRequest, "Invalid id parameter")
		return
	}

	err = h.usecase.DeleteByID(uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.Error(c, http.StatusNotFound, "Data not Found")
			return
		}

		response.Error(c, http.StatusInternalServerError, "failed to delete ErrRecordNotFound")
		return
	}

	response.DeleteSuccess(c, "Work hour deleted successfully")
}
