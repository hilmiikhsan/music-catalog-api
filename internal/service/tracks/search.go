package tracks

import (
	"context"

	"github.com/hilmiikhsan/music-catalog/internal/models/spotify"
	"github.com/hilmiikhsan/music-catalog/internal/models/track_activities"
	spotifyRepo "github.com/hilmiikhsan/music-catalog/internal/repository/spotify"
	"github.com/rs/zerolog/log"
)

func (s *service) Search(ctx context.Context, query string, pageSize, pageIndex int, userID uint) (*spotify.SearchRespose, error) {
	limit := pageSize
	offset := (pageIndex - 1) * pageSize

	trackDetails, err := s.spotifyOutbound.Search(ctx, query, limit, offset)
	if err != nil {
		log.Error().Err(err).Msg("failed to search track")
		return nil, err
	}

	trackIDs := make([]string, len(trackDetails.Tracks.Items))
	for idx, track := range trackDetails.Tracks.Items {
		trackIDs[idx] = track.ID
	}

	trackActivities, err := s.trackActivitiesRepository.GetBulkSpotifyIDs(ctx, userID, trackIDs)
	if err != nil {
		log.Error().Err(err).Msg("failed to get track activities")
		return nil, err
	}

	return modelToResponse(trackDetails, trackActivities), nil
}

func modelToResponse(data *spotifyRepo.SpotifySearchResposes, mapTrackActivities map[string]track_activities.TrackActivity) *spotify.SearchRespose {
	if data == nil {
		return nil
	}

	items := make([]spotify.SpotifyTrackObject, 0)

	for _, item := range data.Tracks.Items {
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

	return &spotify.SearchRespose{
		Items:  items,
		Total:  data.Tracks.Total,
		Limit:  data.Tracks.Limit,
		Offset: data.Tracks.Offset,
	}
}
