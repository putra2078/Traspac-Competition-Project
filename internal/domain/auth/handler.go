package auth

import (
	"net/http"

	"hrm-app/config"
	"hrm-app/internal/domain/user"
	"hrm-app/internal/pkg/utils"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	userRepo user.Repository
	cfg      *config.Config
}

func NewHandler(repo user.Repository, cfg *config.Config) *Handler {
	return &Handler{userRepo: repo, cfg: cfg}
}

type loginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) Login(c *gin.Context) {
	var req loginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "BadRequest",
			"message": err.Error(),
		})
		return
	}

	user, err := h.userRepo.FindByEmail(req.Email)
	if err != nil || user.ID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "Invalid email or password",
		})
		return
	}

	if !utils.CheckPasswordHash(req.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "Invalid email or password",
		})
		return
	}

	token, err := utils.GenerateToken(h.cfg, user.ID, user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "InternalServerError",
			"message": "Failed to generate token",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token": token,
		"token_type":   "Bearer",
		"expires_in":   h.cfg.JWT.ExpiresInMinutes * 60,
	})
}
