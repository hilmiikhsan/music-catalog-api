package tracks

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/hilmiikhsan/music-catalog/internal/models/track_activities"
	"github.com/hilmiikhsan/music-catalog/pkg/jwt"
	"github.com/stretchr/testify/assert"
	gomock "go.uber.org/mock/gomock"
)

func TestHandler_UpsertTrackActivities(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	mockSvc := NewMockservice(ctrlMock)

	isLikeTrue := true

	tests := []struct {
		name               string
		mockFn             func()
		expectedStatusCode int
	}{
		{
			name: "success",
			mockFn: func() {
				mockSvc.EXPECT().UpsertTrackActivities(gomock.Any(), uint(1), track_activities.TrackActivityRequest{
					SpotifyID: "spotifyID",
					IsLike:    &isLikeTrue,
				}).Return(nil)
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			name: "failed",
			mockFn: func() {
				mockSvc.EXPECT().UpsertTrackActivities(gomock.Any(), uint(1), track_activities.TrackActivityRequest{
					SpotifyID: "spotifyID",
					IsLike:    &isLikeTrue,
				}).Return(assert.AnError)
			},
			expectedStatusCode: http.StatusInternalServerError,
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

			endpoint := `/tracks/track-activity`

			payload := track_activities.TrackActivityRequest{
				SpotifyID: "spotifyID",
				IsLike:    &isLikeTrue,
			}

			payloadBytes, err := json.Marshal(payload)
			assert.NoError(t, err)

			req, err := http.NewRequest(http.MethodPost, endpoint, io.NopCloser(bytes.NewBuffer(payloadBytes)))
			assert.NoError(t, err)

			token, err := jwt.CreateToken(1, "username", "")
			assert.NoError(t, err)

			req.Header.Set("Authorization", token)

			h.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatusCode, w.Code)
		})
	}
}
