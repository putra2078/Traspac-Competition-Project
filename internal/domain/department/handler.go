package department

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

func (h *Handler) Register(c *gin.Context) {
	var department Department
	if err := c.ShouldBindJSON(&department); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.usecase.Register(&department); err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Department created successfully"})
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

func (h *Handler) GetBySlug(c *gin.Context) {
	slug := c.Param("slug")
	data, err := h.usecase.GetBySlug(string(slug))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Department not found"})
		return
	}

	c.JSON(http.StatusOK, data)
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
