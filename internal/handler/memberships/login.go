package memberships

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hilmiikhsan/music-catalog/internal/models/memberships"
	"github.com/rs/zerolog/log"
)

func (h *Handler) Login(c *gin.Context) {
	var req memberships.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error().Err(err).Msg("failed to bind request")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	accessToken, err := h.service.Login(req)
	if err != nil {
		log.Error().Err(err).Msg("failed to login")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, memberships.LoginResponse{
		AccessToken: accessToken,
	})
}
