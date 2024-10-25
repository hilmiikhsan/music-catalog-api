package memberships

import (
	"errors"

	"github.com/hilmiikhsan/music-catalog/internal/models/memberships"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func (s *service) SignUp(req memberships.SignUpRequest) error {
	user, err := s.repository.GetUser(req.Email, req.Username, 0)
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Error().Err(err).Msg("failed to get user")
		return err
	}

	if user != nil {
		log.Error().Msg("email or username already exists")
		return errors.New("email or username already exists")
	}

	password, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Error().Err(err).Msg("failed to hash password")
		return err
	}

	model := memberships.User{
		Email:     req.Email,
		Username:  req.Username,
		Password:  string(password),
		CreatedBy: req.Email,
		UpdatedBy: req.Email,
	}

	return s.repository.CreateUser(model)
}
