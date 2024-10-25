package memberships

import (
	"github.com/rs/zerolog/log"

	"github.com/hilmiikhsan/music-catalog/internal/models/memberships"
)

func (r *repository) CreateUser(model memberships.User) error {
	return r.db.Create(&model).Error
}

func (r *repository) GetUser(email, username string, id uint) (*memberships.User, error) {
	user := memberships.User{}

	res := r.db.Where("email = ?", email).Or("username = ?", username).Or("id = ?", id).First(&user)
	if res.Error != nil {
		log.Error().Err(res.Error).Msg("failed to get user")
		return nil, res.Error
	}

	return &user, nil
}
