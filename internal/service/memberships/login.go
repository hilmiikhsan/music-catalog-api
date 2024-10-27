package memberships

import (
	"github.com/hilmiikhsan/music-catalog/internal/models/memberships"
	"github.com/hilmiikhsan/music-catalog/pkg/jwt"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func (s *service) Login(req memberships.LoginRequest) (string, error) {
	user, err := s.repository.GetUser(req.Email, "", 0)
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Error().Err(err).Msg("failed to get user")
		return "", err
	}

	if user == nil {
		log.Error().Msg("email doeant exist")
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		log.Error().Err(err).Msg("failed to compare password")
		return "", err
	}

	accessToken, err := jwt.CreateToken(int64(user.ID), user.Username, s.cfg.Service.SecretKey)
	if err != nil {
		log.Error().Err(err).Msg("failed to create token")
		return "", err
	}

	return accessToken, nil
}
