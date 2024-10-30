package tracks

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/hilmiikhsan/music-catalog/internal/middleware"
	"github.com/hilmiikhsan/music-catalog/internal/models/spotify"
)

//go:generate mockgen -source=handler.go -destination=handler_mock_test.go -package=tracks
type service interface {
	Search(ctx context.Context, query string, pageSize, pageIndex int) (*spotify.SearchRespose, error)
}

type Handler struct {
	*gin.Engine
	service service
}

func NewHandler(api *gin.Engine, service service) *Handler {
	return &Handler{
		Engine:  api,
		service: service,
	}
}

func (h *Handler) RegisterRoute() {
	routes := h.Group("/tracks")
	routes.Use(middleware.AuthMiddleware())

	routes.GET("/search", h.Search)
}
