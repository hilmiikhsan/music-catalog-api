package tracks

import (
	"context"

	"github.com/hilmiikhsan/music-catalog/internal/models/spotify"
	"github.com/hilmiikhsan/music-catalog/internal/models/track_activities"
	spotifyRepo "github.com/hilmiikhsan/music-catalog/internal/repository/spotify"
	"github.com/rs/zerolog/log"
)

func (s *service) GetRecommendation(ctx context.Context, userID uint, limit int, trackID string) (*spotify.RecommendationResponse, error) {
	response, err := s.spotifyOutbound.GetRecommendation(ctx, limit, trackID)
	if err != nil {
		log.Error().Err(err).Msg("failed to get recommendation from spotify")
		return nil, err
	}

	trackIDs := make([]string, len(response.Tracks))
	for idx, track := range response.Tracks {
		trackIDs[idx] = track.ID
	}

	trackActivities, err := s.trackActivitiesRepository.GetBulkSpotifyIDs(ctx, userID, trackIDs)
	if err != nil {
		log.Error().Err(err).Msg("failed to get track activities")
		return nil, err
	}

	return modelToRecommendationResponse(response, trackActivities), nil
}

func modelToRecommendationResponse(data *spotifyRepo.SpotifyRecommendationResponse, mapTrackActivities map[string]track_activities.TrackActivity) *spotify.RecommendationResponse {
	if data == nil {
		return nil
	}

	items := make([]spotify.SpotifyTrackObject, 0)

	for _, item := range data.Tracks {
		artistsName := make([]string, len(item.Artists))
		for idx, artist := range item.Artists {
			artistsName[idx] = artist.Name
		}

		imagesUrl := make([]string, len(item.Album.Images))
		for idx, image := range item.Album.Images {
			imagesUrl[idx] = image.Url
		}

		items = append(items, spotify.SpotifyTrackObject{
			AlbumType:        item.Album.AlbumType,
			AlbumTotalTracks: item.Album.TotalTracks,
			AlbumImagesUrl:   imagesUrl,
			AlbumName:        item.Album.Name,
			ArtistsName:      artistsName,
			Explicit:         item.Explicit,
			ID:               item.ID,
			Name:             item.Name,
			IsLike:           mapTrackActivities[item.ID].IsLike,
		})
	}

	return &spotify.RecommendationResponse{
		Items: items,
	}
}
