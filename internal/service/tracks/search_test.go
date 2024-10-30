package tracks

import (
	"context"
	"reflect"
	"testing"

	"github.com/hilmiikhsan/music-catalog/internal/models/spotify"
	spotyfyOutbound "github.com/hilmiikhsan/music-catalog/internal/repository/spotify"
	"github.com/stretchr/testify/assert"
	gomock "go.uber.org/mock/gomock"
)

func Test_service_Search(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	mockSpotifyOutbound := NewMockspotifyOutbound(ctrlMock)
	next := "https://api.spotify.com/v1/search?query=Bruno+Mars&type=track&market=ID&locale=en-US%2Cen%3Bq%3D0.9&offset=10&limit=10"

	type args struct {
		query     string
		pageSize  int
		pageIndex int
	}
	tests := []struct {
		name    string
		args    args
		want    *spotify.SearchRespose
		wantErr bool
		mockFn  func(args args)
	}{
		{
			name: "success",
			args: args{
				query:     "Bruno Mars",
				pageSize:  10,
				pageIndex: 1,
			},
			want: &spotify.SearchRespose{
				Items: []spotify.SpotifyTrackObject{
					{
						AlbumType:        "single",
						AlbumTotalTracks: 1,
						AlbumImagesUrl: []string{
							"https://i.scdn.co/image/ab67616d0000b273f8c8297efc6022534f1357e1",
							"https://i.scdn.co/image/ab67616d00001e02f8c8297efc6022534f1357e1",
							"https://i.scdn.co/image/ab67616d00004851f8c8297efc6022534f1357e1",
						},
						AlbumName: "APT.",
						ArtistsName: []string{
							"ROSÉ",
							"Bruno Mars",
						},
						Explicit: false,
						ID:       "5vNRhkKd0yEAg8suGBpjeY",
						Name:     "APT.",
					},
					{
						AlbumType:        "single",
						AlbumTotalTracks: 1,
						AlbumImagesUrl: []string{
							"https://i.scdn.co/image/ab67616d0000b27382ea2e9e1858aa012c57cd45",
							"https://i.scdn.co/image/ab67616d00001e0282ea2e9e1858aa012c57cd45",
							"https://i.scdn.co/image/ab67616d0000485182ea2e9e1858aa012c57cd45",
						},
						AlbumName: "Die With A Smile",
						ArtistsName: []string{
							"Lady Gaga",
							"Bruno Mars",
						},
						Explicit: false,
						ID:       "2plbrEY59IikOBgBGLjaoe",
						Name:     "Die With A Smile",
					},
				},
				Total:  894,
				Limit:  10,
				Offset: 1,
			},
			wantErr: false,
			mockFn: func(args args) {
				mockSpotifyOutbound.EXPECT().Search(gomock.Any(), args.query, 10, 0).Return(&spotyfyOutbound.SpotifySearchResposes{
					Tracks: spotyfyOutbound.SpotifyTracks{
						Href:   "https://api.spotify.com/v1/search?query=Bruno+Mars&type=track&market=ID&locale=en-US%2Cen%3Bq%3D0.9&offset=0&limit=10",
						Limit:  10,
						Next:   &next,
						Offset: 1,
						Total:  894,
						Items: []spotyfyOutbound.SpotifyTrackObject{
							{
								Album: spotyfyOutbound.SpotifyAlbumObject{
									AlbumType:   "single",
									TotalTracks: 1,
									Images: []spotyfyOutbound.SpotifyAlbumImagesObject{
										{Url: "https://i.scdn.co/image/ab67616d0000b273f8c8297efc6022534f1357e1"},
										{Url: "https://i.scdn.co/image/ab67616d00001e02f8c8297efc6022534f1357e1"},
										{Url: "https://i.scdn.co/image/ab67616d00004851f8c8297efc6022534f1357e1"},
									},
									Name: "APT.",
								},
								Artists: []spotyfyOutbound.SpotifyArtistsObject{
									{Name: "ROSÉ"},
									{Name: "Bruno Mars"},
								},
								Explicit: false,
								ID:       "5vNRhkKd0yEAg8suGBpjeY",
								Name:     "APT.",
							},
							{
								Album: spotyfyOutbound.SpotifyAlbumObject{
									AlbumType:   "single",
									TotalTracks: 1,
									Images: []spotyfyOutbound.SpotifyAlbumImagesObject{
										{Url: "https://i.scdn.co/image/ab67616d0000b27382ea2e9e1858aa012c57cd45"},
										{Url: "https://i.scdn.co/image/ab67616d00001e0282ea2e9e1858aa012c57cd45"},
										{Url: "https://i.scdn.co/image/ab67616d0000485182ea2e9e1858aa012c57cd45"},
									},
									Name: "Die With A Smile",
								},
								Artists: []spotyfyOutbound.SpotifyArtistsObject{
									{Name: "Lady Gaga"},
									{Name: "Bruno Mars"},
								},
								Explicit: false,
								ID:       "2plbrEY59IikOBgBGLjaoe",
								Name:     "Die With A Smile",
							},
						},
					},
				}, nil)
			},
		},
		{
			name: "failed",
			args: args{
				query:     "Bruno Mars",
				pageSize:  10,
				pageIndex: 1,
			},
			want:    nil,
			wantErr: true,
			mockFn: func(args args) {
				mockSpotifyOutbound.EXPECT().Search(gomock.Any(), args.query, 10, 0).Return(nil, assert.AnError)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)
			s := &service{
				spotifyOutbound: mockSpotifyOutbound,
			}
			got, err := s.Search(context.Background(), tt.args.query, tt.args.pageSize, tt.args.pageIndex)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.Search() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.Search() = %v, want %v", got, tt.want)
			}
		})
	}
}
