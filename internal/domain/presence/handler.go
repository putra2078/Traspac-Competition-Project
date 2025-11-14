package presence

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"hrm-app/internal/response"
)

type Handler struct {
	usecase UseCase
}

func NewHandler(u UseCase) *Handler {
	return &Handler{usecase: u}
}

func (h *Handler) Checkin(c *gin.Context) {
	// Ambil user_id dari context (dari JWT middleware)
	userID, exists := c.Get("user_id")
	if !exists {
		response.Error(c, http.StatusUnauthorized, "unauthorized: user_id tidak ditemukan di token")
		return
	}

	// Ambil data lokasi dan status dari body JSON
	var req struct {
		LatCheckIn  float64 `json:"lat_check_in"`
		LongCheckIn float64 `json:"long_check_in"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "invalid request body")
		return
	}

	// Jalankan usecase checkin
	err := h.usecase.Checkin(userID.(uint), req.LatCheckIn, req.LongCheckIn)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	response.Success(c, "Check-in berhasil")
}

func (h *Handler) Checkout(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		response.Error(c, http.StatusUnauthorized, "unauthorized: user_id tidak ditemukan di token")
		return
	}

	var req struct {
		LatCheckOut  float64 `json:"lat_check_out"`
		LongCheckOut float64 `json:"long_check_out"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "invalid request body")
		return
	}

	err := h.usecase.Checkout(userID.(uint), req.LatCheckOut, req.LongCheckOut)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	response.Success(c, "Check-out berhasil")
}
