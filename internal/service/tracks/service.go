package tracks

import (
	"context"

	"github.com/hilmiikhsan/music-catalog/internal/models/track_activities"
	"github.com/hilmiikhsan/music-catalog/internal/repository/spotify"
)

//go:generate mockgen -source=service.go -destination=service_mock_test.go -package=tracks
type spotifyOutbound interface {
	Search(ctx context.Context, query string, limit, offset int) (*spotify.SpotifySearchResposes, error)
}

type trackActivitiesRepository interface {
	Create(ctx context.Context, model track_activities.TrackActivity) error
	Update(ctx context.Context, model track_activities.TrackActivity) error
	Get(ctx context.Context, userID uint, spotifyID string) (*track_activities.TrackActivity, error)
	GetBulkSpotifyIDs(ctx context.Context, userID uint, spotifyIDs []string) (map[string]track_activities.TrackActivity, error)
}

type service struct {
	spotifyOutbound           spotifyOutbound
	trackActivitiesRepository trackActivitiesRepository
}

func NewService(spotifyOutbound spotifyOutbound, trackActivitiesRepository trackActivitiesRepository) *service {
	return &service{
		spotifyOutbound:           spotifyOutbound,
		trackActivitiesRepository: trackActivitiesRepository,
	}
}
