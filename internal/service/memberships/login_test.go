package memberships

import (
	"testing"

	"github.com/hilmiikhsan/music-catalog/internal/configs"
	"github.com/hilmiikhsan/music-catalog/internal/models/memberships"
	"github.com/stretchr/testify/assert"
	gomock "go.uber.org/mock/gomock"
	"gorm.io/gorm"
)

func Test_service_Login(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	mockRepo := NewMockrepository(ctrlMock)

	type args struct {
		req memberships.LoginRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		mockFn  func(args args)
	}{
		{
			name: "success",
			args: args{
				req: memberships.LoginRequest{
					Email:    "test@gmail.com",
					Password: "password",
				},
			},
			wantErr: false,
			mockFn: func(args args) {
				mockRepo.EXPECT().GetUser(args.req.Email, "", uint(0)).Return(&memberships.User{
					Model: gorm.Model{
						ID: 1,
					},
					Email:    "test@gmail.com",
					Password: "$2b$12$K8rfJ0vyajqfUH/O6s0E8.muni.uD1bldLfikJ8RCYkMftEuVELMa",
					Username: "testinguser",
				}, nil)
			},
		},
		{
			name: "failed when GetUser",
			args: args{
				req: memberships.LoginRequest{
					Email:    "test@gmail.com",
					Password: "password",
				},
			},
			wantErr: true,
			mockFn: func(args args) {
				mockRepo.EXPECT().GetUser(args.req.Email, "", uint(0)).Return(nil, assert.AnError)
			},
		},
		{
			name: "failed when password not match",
			args: args{
				req: memberships.LoginRequest{
					Email:    "test@gmail.com",
					Password: "password",
				},
			},
			wantErr: true,
			mockFn: func(args args) {
				mockRepo.EXPECT().GetUser(args.req.Email, "", uint(0)).Return(&memberships.User{
					Model: gorm.Model{
						ID: 1,
					},
					Email:    "test@gmail.com",
					Password: "password",
					Username: "testinguser",
				}, nil)
			},
		},
	}
	for _, tt := range tests {
		tt.mockFn(tt.args)
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				cfg: &configs.Config{
					Service: configs.Service{
						SecretKey: "secret",
					},
				},
				repository: mockRepo,
			}
			got, err := s.Login(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				assert.NotEmpty(t, got)
			} else {
				assert.Empty(t, got)
			}
		})
	}
}
