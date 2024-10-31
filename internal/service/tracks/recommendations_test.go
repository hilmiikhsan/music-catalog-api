package tracks

import (
	"context"
	"reflect"
	"testing"

	"github.com/hilmiikhsan/music-catalog/internal/models/spotify"
	"github.com/hilmiikhsan/music-catalog/internal/models/track_activities"
	spotifyRepo "github.com/hilmiikhsan/music-catalog/internal/repository/spotify"
	"github.com/stretchr/testify/assert"
	gomock "go.uber.org/mock/gomock"
)

func Test_service_GetRecommendation(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	mockSpotifyOutbound := NewMockspotifyOutbound(ctrlMock)
	mockTrackActivitiesRepository := NewMocktrackActivitiesRepository(ctrlMock)

	isLikeTrue := true
	isLikeFalse := false

	type args struct {
		userID  uint
		limit   int
		trackID string
	}
	tests := []struct {
		name    string
		args    args
		want    *spotify.RecommendationResponse
		wantErr bool
		mockFn  func(args args)
	}{
		{
			name: "success",
			args: args{
				userID:  1,
				limit:   10,
				trackID: "track_id",
			},
			want: &spotify.RecommendationResponse{
				Items: []spotify.SpotifyTrackObject{
					{
						AlbumType:        "album",
						AlbumTotalTracks: 22,
						AlbumImagesUrl:   []string{"https://i.scdn.co/image/ab67616d0000b273e8b066f70c206551210d902b", "https://i.scdn.co/image/ab67616d00001e02e8b066f70c206551210d902b", "https://i.scdn.co/image/ab67616d00004851e8b066f70c206551210d902b"},
						AlbumName:        "Bohemian Rhapsody (The Original Soundtrack)",
						ArtistsName:      []string{"Queen"},
						Explicit:         false,
						ID:               "3z8h0TU7ReDPLIbEnYhWZb",
						Name:             "Bohemian Rhapsody",
						IsLike:           &isLikeTrue,
					},
					{
						AlbumType:        "album",
						AlbumTotalTracks: 12,
						AlbumImagesUrl:   []string{"https://i.scdn.co/image/ab67616d0000b273e319baafd16e84f0408af2a0", "https://i.scdn.co/image/ab67616d00001e02e319baafd16e84f0408af2a0", "https://i.scdn.co/image/ab67616d00004851e319baafd16e84f0408af2a0"},
						AlbumName:        "A Night At The Opera (2011 Remaster)",
						ArtistsName:      []string{"Queen"},
						Explicit:         false,
						ID:               "4u7EnebtmKWzUH433cf5Qv",
						Name:             "Bohemian Rhapsody - Remastered 2011",
						IsLike:           &isLikeFalse,
					},
				},
			},
			wantErr: false,
			mockFn: func(args args) {
				mockSpotifyOutbound.EXPECT().GetRecommendation(gomock.Any(), 10, "track_id").Return(&spotifyRepo.SpotifyRecommendationResponse{
					Tracks: []spotifyRepo.SpotifyTrackObject{
						{
							Album: spotifyRepo.SpotifyAlbumObject{
								AlbumType:   "album",
								TotalTracks: 22,
								Images: []spotifyRepo.SpotifyAlbumImagesObject{
									{
										Url: "https://i.scdn.co/image/ab67616d0000b273e8b066f70c206551210d902b",
									},
									{
										Url: "https://i.scdn.co/image/ab67616d00001e02e8b066f70c206551210d902b",
									},
									{
										Url: "https://i.scdn.co/image/ab67616d00004851e8b066f70c206551210d902b",
									},
								},
								Name: "Bohemian Rhapsody (The Original Soundtrack)",
							},
							Artists: []spotifyRepo.SpotifyArtistsObject{
								{
									Href: "https://api.spotify.com/v1/artists/1dfeR4HaWDbWqFHLkxsg1d",
									Name: "Queen",
								},
							},
							Explicit: false,
							Href:     "https://api.spotify.com/v1/tracks/3z8h0TU7ReDPLIbEnYhWZb",
							ID:       "3z8h0TU7ReDPLIbEnYhWZb",
							Name:     "Bohemian Rhapsody",
						},
						{
							Album: spotifyRepo.SpotifyAlbumObject{
								AlbumType:   "album",
								TotalTracks: 12,
								Images: []spotifyRepo.SpotifyAlbumImagesObject{
									{
										Url: "https://i.scdn.co/image/ab67616d0000b273e319baafd16e84f0408af2a0",
									},
									{
										Url: "https://i.scdn.co/image/ab67616d00001e02e319baafd16e84f0408af2a0",
									},
									{
										Url: "https://i.scdn.co/image/ab67616d00004851e319baafd16e84f0408af2a0",
									},
								},
								Name: "A Night At The Opera (2011 Remaster)",
							},
							Artists: []spotifyRepo.SpotifyArtistsObject{
								{
									Href: "https://api.spotify.com/v1/artists/1dfeR4HaWDbWqFHLkxsg1d",
									Name: "Queen",
								},
							},
							Explicit: false,
							Href:     "https://api.spotify.com/v1/tracks/4u7EnebtmKWzUH433cf5Qv",
							ID:       "4u7EnebtmKWzUH433cf5Qv",
							Name:     "Bohemian Rhapsody - Remastered 2011",
						},
					},
				}, nil)

				mockTrackActivitiesRepository.EXPECT().GetBulkSpotifyIDs(gomock.Any(), uint(1), []string{
					"3z8h0TU7ReDPLIbEnYhWZb",
					"4u7EnebtmKWzUH433cf5Qv",
				}).
					Return(map[string]track_activities.TrackActivity{
						"3z8h0TU7ReDPLIbEnYhWZb": {
							IsLike: &isLikeTrue,
						},
						"4u7EnebtmKWzUH433cf5Qv": {
							IsLike: &isLikeFalse,
						},
					}, nil)
			},
		},
		{
			name: "failed when get recommendation from spotify outbound",
			args: args{
				userID:  1,
				limit:   10,
				trackID: "trackID",
			},
			want:    nil,
			wantErr: true,
			mockFn: func(args args) {
				mockSpotifyOutbound.EXPECT().GetRecommendation(gomock.Any(), 10, "trackID").Return(nil, assert.AnError)
			},
		},
		{
			name: "failed when get bulk spotify id",
			args: args{
				userID:  1,
				limit:   10,
				trackID: "trackID",
			},
			want:    nil,
			wantErr: true,
			mockFn: func(args args) {
				mockSpotifyOutbound.EXPECT().GetRecommendation(gomock.Any(), 10, "trackID").Return(&spotifyRepo.SpotifyRecommendationResponse{
					Tracks: []spotifyRepo.SpotifyTrackObject{
						{
							Album: spotifyRepo.SpotifyAlbumObject{
								AlbumType:   "album",
								TotalTracks: 22,
								Images: []spotifyRepo.SpotifyAlbumImagesObject{
									{
										Url: "https://i.scdn.co/image/ab67616d0000b273e8b066f70c206551210d902b",
									},
									{
										Url: "https://i.scdn.co/image/ab67616d00001e02e8b066f70c206551210d902b",
									},
									{
										Url: "https://i.scdn.co/image/ab67616d00004851e8b066f70c206551210d902b",
									},
								},
								Name: "Bohemian Rhapsody (The Original Soundtrack)",
							},
							Artists: []spotifyRepo.SpotifyArtistsObject{
								{
									Href: "https://api.spotify.com/v1/artists/1dfeR4HaWDbWqFHLkxsg1d",
									Name: "Queen",
								},
							},
							Explicit: false,
							Href:     "https://api.spotify.com/v1/tracks/3z8h0TU7ReDPLIbEnYhWZb",
							ID:       "3z8h0TU7ReDPLIbEnYhWZb",
							Name:     "Bohemian Rhapsody",
						},
						{
							Album: spotifyRepo.SpotifyAlbumObject{
								AlbumType:   "album",
								TotalTracks: 12,
								Images: []spotifyRepo.SpotifyAlbumImagesObject{
									{
										Url: "https://i.scdn.co/image/ab67616d0000b273e319baafd16e84f0408af2a0",
									},
									{
										Url: "https://i.scdn.co/image/ab67616d00001e02e319baafd16e84f0408af2a0",
									},
									{
										Url: "https://i.scdn.co/image/ab67616d00004851e319baafd16e84f0408af2a0",
									},
								},
								Name: "A Night At The Opera (2011 Remaster)",
							},
							Artists: []spotifyRepo.SpotifyArtistsObject{
								{
									Href: "https://api.spotify.com/v1/artists/1dfeR4HaWDbWqFHLkxsg1d",
									Name: "Queen",
								},
							},
							Explicit: false,
							Href:     "https://api.spotify.com/v1/tracks/4u7EnebtmKWzUH433cf5Qv",
							ID:       "4u7EnebtmKWzUH433cf5Qv",
							Name:     "Bohemian Rhapsody - Remastered 2011",
						},
					},
				}, nil)

				mockTrackActivitiesRepository.EXPECT().GetBulkSpotifyIDs(gomock.Any(), uint(1), []string{
					"3z8h0TU7ReDPLIbEnYhWZb",
					"4u7EnebtmKWzUH433cf5Qv",
				}).Return(nil, assert.AnError)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)
			s := &service{
				spotifyOutbound:           mockSpotifyOutbound,
				trackActivitiesRepository: mockTrackActivitiesRepository,
			}
			got, err := s.GetRecommendation(context.Background(), tt.args.userID, tt.args.limit, tt.args.trackID)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.GetRecommendation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.GetRecommendation() = %v, want %v", got, tt.want)
			}
		})
	}
}
