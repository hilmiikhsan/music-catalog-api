package memberships

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/hilmiikhsan/music-catalog/internal/models/memberships"
	"github.com/stretchr/testify/assert"
	gomock "go.uber.org/mock/gomock"
)

func TestHandler_Login(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	mockSvc := NewMockservice(ctrlMock)

	tests := []struct {
		name               string
		mockFn             func()
		expectedStatusCode int
		expectedBody       memberships.LoginResponse
		wantErr            bool
	}{
		{
			name: "success",
			mockFn: func() {
				mockSvc.EXPECT().Login(memberships.LoginRequest{
					Email:    "test@gmail.com",
					Password: "password",
				}).Return("access_token", nil)
			},
			expectedStatusCode: http.StatusOK,
			expectedBody: memberships.LoginResponse{
				AccessToken: "access_token",
			},
			wantErr: false,
		},
		{
			name: "failed",
			mockFn: func() {
				mockSvc.EXPECT().Login(memberships.LoginRequest{
					Email:    "test@gmail.com",
					Password: "password",
				}).Return("", assert.AnError)
			},
			expectedStatusCode: http.StatusInternalServerError,
			expectedBody:       memberships.LoginResponse{},
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

			endpoint := `/memberships/login`

			model := memberships.LoginRequest{
				Email:    "test@gmail.com",
				Password: "password",
			}

			val, err := json.Marshal(model)
			assert.NoError(t, err)

			body := bytes.NewReader(val)

			req, err := http.NewRequest(http.MethodPost, endpoint, body)
			assert.NoError(t, err)

			h.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatusCode, w.Code)

			if !tt.wantErr {
				res := w.Result()
				defer res.Body.Close()

				response := memberships.LoginResponse{}
				err = json.NewDecoder(res.Body).Decode(&response)
				assert.NoError(t, err)

				assert.Equal(t, tt.expectedBody, response)
			}
		})
	}
}
