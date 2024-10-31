package tracks

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hilmiikhsan/music-catalog/internal/models/track_activities"
	"github.com/rs/zerolog/log"
)

func (h *Handler) UpsertTrackActivities(c *gin.Context) {
	ctx := c.Request.Context()

	var req track_activities.TrackActivityRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error().Err(err).Msg("failed to bind request")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := c.GetUint("user_id")

	err := h.service.UpsertTrackActivities(ctx, userID, req)
	if err != nil {
		log.Error().Err(err).Msg("failed to upsert track activities")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}
