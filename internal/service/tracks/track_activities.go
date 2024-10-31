package tracks

import (
	"context"
	"fmt"
	"log"

	"github.com/hilmiikhsan/music-catalog/internal/models/track_activities"
	"gorm.io/gorm"
)

func (s *service) UpsertTrackActivities(ctx context.Context, userID uint, req track_activities.TrackActivityRequest) error {
	activity, err := s.trackActivitiesRepository.Get(ctx, userID, req.SpotifyID)
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Println("failed to get track activity: ", err)
		return err
	}

	if err == gorm.ErrRecordNotFound || activity == nil {
		err = s.trackActivitiesRepository.Create(ctx, track_activities.TrackActivity{
			UserID:    userID,
			SpotifyID: req.SpotifyID,
			IsLike:    req.IsLike,
			CreatedBy: fmt.Sprintf("%d", userID),
			UpdatedBy: fmt.Sprintf("%d", userID),
		})
		if err != nil {
			log.Println("failed to create track activity: ", err)
			return err
		}

		return nil
	}

	activity.IsLike = req.IsLike

	err = s.trackActivitiesRepository.Update(ctx, *activity)
	if err != nil {
		log.Println("failed to update track activity: ", err)
		return err
	}

	return nil
}
