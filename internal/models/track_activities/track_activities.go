package track_activities

import "gorm.io/gorm"

type (
	TrackActivity struct {
		gorm.Model
		UserID    uint   `gorm:"not null"`
		SpotifyID string `gorm:"not null"`
		IsLike    *bool  `gorm:"null"`
		CreatedBy string `gorm:"not null"`
		UpdatedBy string `gorm:"not null"`
	}
)

type (
	TrackActivityRequest struct {
		SpotifyID string `json:"spotify_id"`
		IsLike    *bool  `json:"is_like"`
	}
)
