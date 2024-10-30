package spotify

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"testing"
	"time"

	"github.com/hilmiikhsan/music-catalog/internal/configs"
	"github.com/hilmiikhsan/music-catalog/pkg/httpclient"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func Test_outbound_Search(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	mockHTTPClient := httpclient.NewMockHTTPClient(ctrlMock)
	next := "https://api.spotify.com/v1/search?query=Bruno+Mars&type=track&market=ID&locale=en-US%2Cen%3Bq%3D0.9&offset=10&limit=10"

	type args struct {
		query  string
		limit  int
		offset int
	}
	tests := []struct {
		name    string
		args    args
		want    *SpotifySearchResposes
		wantErr bool
		mockFn  func(args args)
	}{
		{
			name: "success",
			args: args{
				query:  "Bruno Mars",
				limit:  10,
				offset: 0,
			},
			want: &SpotifySearchResposes{
				Tracks: SpotifyTracks{
					Href:   "https://api.spotify.com/v1/search?query=Bruno+Mars&type=track&market=ID&locale=en-US%2Cen%3Bq%3D0.9&offset=0&limit=10",
					Limit:  10,
					Next:   &next,
					Offset: 0,
					Total:  894,
					Items: []SpotifyTrackObject{
						{
							Album: SpotifyAlbumObject{
								AlbumType:   "single",
								TotalTracks: 1,
								Images: []SpotifyAlbumImagesObject{
									{
										Url: "https://i.scdn.co/image/ab67616d0000b273f8c8297efc6022534f1357e1",
									},
									{
										Url: "https://i.scdn.co/image/ab67616d00001e02f8c8297efc6022534f1357e1",
									},
									{
										Url: "https://i.scdn.co/image/ab67616d00004851f8c8297efc6022534f1357e1",
									},
								},
								Name: "APT.",
							},
							Artists: []SpotifyArtistsObject{
								{
									Href: "https://api.spotify.com/v1/artists/3eVa5w3URK5duf6eyVDbu9",
									Name: "ROSÉ",
								},
								{
									Href: "https://api.spotify.com/v1/artists/0du5cEVh5yTK9QJze8zA0C",
									Name: "Bruno Mars",
								},
							},
							Explicit: false,
							Href:     "https://api.spotify.com/v1/tracks/5vNRhkKd0yEAg8suGBpjeY",
							ID:       "5vNRhkKd0yEAg8suGBpjeY",
							Name:     "APT.",
						},
						{
							Album: SpotifyAlbumObject{
								AlbumType:   "single",
								TotalTracks: 1,
								Images: []SpotifyAlbumImagesObject{
									{
										Url: "https://i.scdn.co/image/ab67616d0000b27382ea2e9e1858aa012c57cd45",
									},
									{
										Url: "https://i.scdn.co/image/ab67616d00001e0282ea2e9e1858aa012c57cd45",
									},
									{
										Url: "https://i.scdn.co/image/ab67616d0000485182ea2e9e1858aa012c57cd45",
									},
								},
								Name: "Die With A Smile",
							},
							Artists: []SpotifyArtistsObject{
								{
									Href: "https://api.spotify.com/v1/artists/1HY2Jd0NmPuamShAr6KMms",
									Name: "Lady Gaga",
								},
								{
									Href: "https://api.spotify.com/v1/artists/0du5cEVh5yTK9QJze8zA0C",
									Name: "Bruno Mars",
								},
							},
							Explicit: false,
							Href:     "https://api.spotify.com/v1/tracks/2plbrEY59IikOBgBGLjaoe",
							ID:       "2plbrEY59IikOBgBGLjaoe",
							Name:     "Die With A Smile",
						},
					},
				},
			},
			wantErr: false,
			mockFn: func(args args) {
				params := url.Values{}

				params.Set("q", args.query)
				params.Set("type", "track")
				params.Set("limit", strconv.Itoa(args.limit))
				params.Set("offset", strconv.Itoa(args.offset))

				basePath := "https://api.spotify.com/v1/search"
				urlPath := fmt.Sprintf("%s?%s", basePath, params.Encode())

				req, err := http.NewRequest(http.MethodGet, urlPath, nil)
				assert.NoError(t, err)

				req.Header.Set("Authorization", "Bearer access_token")

				mockHTTPClient.EXPECT().Do(req).Return(&http.Response{
					StatusCode: http.StatusOK,
					Body:       io.NopCloser(bytes.NewBufferString(searchResponse)),
				}, nil)
			},
		},
		{
			name: "failed",
			args: args{
				query:  "Bruno Mars",
				limit:  10,
				offset: 0,
			},
			want:    nil,
			wantErr: true,
			mockFn: func(args args) {
				params := url.Values{}

				params.Set("q", args.query)
				params.Set("type", "track")
				params.Set("limit", strconv.Itoa(args.limit))
				params.Set("offset", strconv.Itoa(args.offset))

				basePath := "https://api.spotify.com/v1/search"
				urlPath := fmt.Sprintf("%s?%s", basePath, params.Encode())

				req, err := http.NewRequest(http.MethodGet, urlPath, nil)
				assert.NoError(t, err)

				req.Header.Set("Authorization", "Bearer access_token")

				mockHTTPClient.EXPECT().Do(req).Return(&http.Response{
					StatusCode: http.StatusInternalServerError,
					Body:       io.NopCloser(bytes.NewBufferString(`Internal Server Error`)),
				}, nil)
			},
		},
	}
	for _, tt := range tests {
		tt.mockFn(tt.args)
		t.Run(tt.name, func(t *testing.T) {
			o := &outbound{
				cfg:         &configs.Config{},
				client:      mockHTTPClient,
				AccessToken: "access_token",
				TokenType:   "Bearer",
				ExpiredAt:   time.Now().Add(1 * time.Hour),
			}
			got, err := o.Search(context.Background(), tt.args.query, tt.args.limit, tt.args.offset)
			if (err != nil) != tt.wantErr {
				t.Errorf("outbound.Search() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("outbound.Search() = %v, want %v", got, tt.want)
			}
		})
	}
}
