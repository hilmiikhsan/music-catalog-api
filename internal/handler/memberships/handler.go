package memberships

import (
	"github.com/gin-gonic/gin"
	"github.com/hilmiikhsan/music-catalog/internal/models/memberships"
)

//go:generate mockgen -source=handler.go -destination=handler_mock_test.go -package=memberships
type service interface {
	SignUp(req memberships.SignUpRequest) error
	Login(req memberships.LoginRequest) (string, error)
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
	routes := h.Group("/memberships")
	routes.POST("/signup", h.SignUp)
	routes.POST("/login", h.Login)
}
