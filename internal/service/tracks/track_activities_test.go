package tracks

import (
	"context"
	"fmt"
	"testing"

	"github.com/hilmiikhsan/music-catalog/internal/models/track_activities"
	"github.com/stretchr/testify/assert"
	gomock "go.uber.org/mock/gomock"
	"gorm.io/gorm"
)

func Test_service_UpsertTrackActivities(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	mockTrackActivitiesRepo := NewMocktrackActivitiesRepository(ctrlMock)

	isLikeTrue := true
	isLikeFalse := false

	type args struct {
		userID uint
		req    track_activities.TrackActivityRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		mockFn  func(args args)
	}{
		{
			name: "success create",
			args: args{
				userID: 1,
				req: track_activities.TrackActivityRequest{
					SpotifyID: "spotifyID",
					IsLike:    &isLikeTrue,
				},
			},
			wantErr: false,
			mockFn: func(args args) {
				mockTrackActivitiesRepo.EXPECT().Get(gomock.Any(), args.userID, args.req.SpotifyID).Return(nil, gorm.ErrRecordNotFound)
				mockTrackActivitiesRepo.EXPECT().Create(gomock.Any(), track_activities.TrackActivity{
					UserID:    args.userID,
					SpotifyID: args.req.SpotifyID,
					IsLike:    args.req.IsLike,
					CreatedBy: fmt.Sprintf("%d", args.userID),
					UpdatedBy: fmt.Sprintf("%d", args.userID),
				}).Return(nil)
			},
		},
		{
			name: "success update",
			args: args{
				userID: 1,
				req: track_activities.TrackActivityRequest{
					SpotifyID: "spotifyID",
					IsLike:    &isLikeTrue,
				},
			},
			wantErr: false,
			mockFn: func(args args) {
				mockTrackActivitiesRepo.EXPECT().Get(gomock.Any(), args.userID, args.req.SpotifyID).Return(&track_activities.TrackActivity{
					IsLike: &isLikeFalse,
				}, nil)
				mockTrackActivitiesRepo.EXPECT().Update(gomock.Any(), track_activities.TrackActivity{
					IsLike: args.req.IsLike,
				}).Return(nil)
			},
		},
		{
			name: "failed",
			args: args{
				userID: 1,
				req: track_activities.TrackActivityRequest{
					SpotifyID: "spotifyID",
					IsLike:    &isLikeTrue,
				},
			},
			wantErr: true,
			mockFn: func(args args) {
				mockTrackActivitiesRepo.EXPECT().Get(gomock.Any(), args.userID, args.req.SpotifyID).Return(nil, assert.AnError)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)
			s := &service{
				trackActivitiesRepository: mockTrackActivitiesRepo,
			}
			if err := s.UpsertTrackActivities(context.Background(), tt.args.userID, tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("service.UpsertTrackActivities() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
