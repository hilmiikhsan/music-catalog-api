package memberships

import (
	"testing"

	"github.com/hilmiikhsan/music-catalog/internal/configs"
	"github.com/hilmiikhsan/music-catalog/internal/models/memberships"
	"github.com/stretchr/testify/assert"
	gomock "go.uber.org/mock/gomock"
	"gorm.io/gorm"
)

func Test_service_SignUp(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	mockRepo := NewMockrepository(ctrlMock)

	type args struct {
		req memberships.SignUpRequest
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
				req: memberships.SignUpRequest{
					Email:    "test@gmail.com",
					Username: "testinguser",
					Password: "password",
				},
			},
			wantErr: false,
			mockFn: func(args args) {
				mockRepo.EXPECT().GetUser(args.req.Email, args.req.Username, uint(0)).Return(nil, gorm.ErrRecordNotFound)
				mockRepo.EXPECT().CreateUser(gomock.Any()).Return(nil)
			},
		},
		{
			name: "failed when GetUser",
			args: args{
				req: memberships.SignUpRequest{
					Email:    "test@gmail.com",
					Username: "testinguser",
					Password: "password",
				},
			},
			wantErr: true,
			mockFn: func(args args) {
				mockRepo.EXPECT().GetUser(args.req.Email, args.req.Username, uint(0)).Return(nil, assert.AnError)
			},
		},
		{
			name: "failed when user already exists",
			args: args{
				req: memberships.SignUpRequest{
					Email:    "existinguser@gmail.com",
					Username: "existinguser",
					Password: "password",
				},
			},
			wantErr: true,
			mockFn: func(args args) {
				mockRepo.EXPECT().GetUser(args.req.Email, args.req.Username, uint(0)).Return(&memberships.User{}, nil)
			},
		},
		{
			name: "failed when CreateUser",
			args: args{
				req: memberships.SignUpRequest{
					Email:    "test@gmail.com",
					Username: "testinguser",
					Password: "password",
				},
			},
			wantErr: true,
			mockFn: func(args args) {
				mockRepo.EXPECT().GetUser(args.req.Email, args.req.Username, uint(0)).Return(nil, gorm.ErrRecordNotFound)
				mockRepo.EXPECT().CreateUser(gomock.Any()).Return(assert.AnError)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)
			s := &service{
				cfg:        &configs.Config{},
				repository: mockRepo,
			}
			if err := s.SignUp(tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("service.SignUp() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
