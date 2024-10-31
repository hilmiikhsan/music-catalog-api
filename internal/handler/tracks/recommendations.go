package tracks

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetRecommendation(c *gin.Context) {
	ctx := c.Request.Context()

	trackID := c.Query("track_id")
	limitStr := c.Query("limit")

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = 10
	}

	userID := c.GetUint("user_id")

	response, err := h.service.GetRecommendation(ctx, userID, limit, trackID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response)
}
