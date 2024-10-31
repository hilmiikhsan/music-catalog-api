package tracks

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/hilmiikhsan/music-catalog/internal/models/spotify"
	"github.com/hilmiikhsan/music-catalog/pkg/jwt"
	"github.com/stretchr/testify/assert"
	gomock "go.uber.org/mock/gomock"
)

func TestHandler_Search(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	mockSvc := NewMockservice(ctrlMock)

	tests := []struct {
		name               string
		mockFn             func()
		expectedStatusCode int
		expectedBody       spotify.SearchRespose
		wantErr            bool
	}{
		{
			name: "success",
			mockFn: func() {
				mockSvc.EXPECT().Search(gomock.Any(), "Bruno Mars", 10, 1, uint(1)).Return(&spotify.SearchRespose{
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
				}, nil)
			},
			expectedStatusCode: http.StatusOK,
			expectedBody: spotify.SearchRespose{
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
		},
		{
			name: "failed",
			mockFn: func() {
				mockSvc.EXPECT().Search(gomock.Any(), "Bruno Mars", 10, 1, uint(1)).Return(nil, assert.AnError)
			},
			expectedStatusCode: http.StatusInternalServerError,
			expectedBody:       spotify.SearchRespose{},
			wantErr:            true,
		},
	}
	for _, tt := range tests {
		tt.mockFn()
		api := gin.New()

		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{
				Engine:  api,
				service: mockSvc,
			}

			h.RegisterRoute()

			w := httptest.NewRecorder()

			endpoint := `/tracks/search?query=Bruno+Mars&pageSize=10&pageIndex=1`

			req, err := http.NewRequest(http.MethodGet, endpoint, nil)
			assert.NoError(t, err)

			token, err := jwt.CreateToken(1, "username", "")
			assert.NoError(t, err)

			req.Header.Set("Authorization", token)

			h.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatusCode, w.Code)

			if !tt.wantErr {
				res := w.Result()
				defer res.Body.Close()

				response := spotify.SearchRespose{}
				err = json.NewDecoder(res.Body).Decode(&response)
				assert.NoError(t, err)

				assert.Equal(t, tt.expectedBody, response)
			}
		})
	}
}
