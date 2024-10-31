package track_activities

import (
	"context"

	"github.com/hilmiikhsan/music-catalog/internal/models/track_activities"
	"github.com/rs/zerolog/log"
)

func (r *repository) Create(ctx context.Context, model track_activities.TrackActivity) error {
	return r.db.Create(&model).Error
}

func (r *repository) Update(ctx context.Context, model track_activities.TrackActivity) error {
	return r.db.Save(&model).Error
}

func (r *repository) Get(ctx context.Context, userID uint, spotifyID string) (*track_activities.TrackActivity, error) {
	activity := track_activities.TrackActivity{}

	res := r.db.Where("user_id = ?", userID).Where("spotify_id = ?", spotifyID).First(&activity)
	if res.Error != nil {
		log.Error().Err(res.Error).Msg("failed to get track activity")
		return nil, res.Error
	}

	return &activity, nil
}

func (r *repository) GetBulkSpotifyIDs(ctx context.Context, userID uint, spotifyIDs []string) (map[string]track_activities.TrackActivity, error) {
	activities := make([]track_activities.TrackActivity, 0)

	res := r.db.Where("user_id = ?", userID).Where("spotify_id IN ?", spotifyIDs).First(&activities)
	if res.Error != nil {
		log.Error().Err(res.Error).Msg("failed to get track activity")
		return nil, res.Error
	}

	result := make(map[string]track_activities.TrackActivity, 0)
	for _, activity := range activities {
		result[activity.SpotifyID] = activity
	}

	return result, nil
}
