package usecase

import (
	"context"
	"errors"
	"myapp/internal/domain/model"
	"myapp/internal/domain/repository"
	"myapp/internal/domain/repository/repository_mock"
	"myapp/internal/pkg/hash"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestNewUserUsecase(t *testing.T) {
	type args struct {
		r repository.UserRepository
	}
	tests := []struct {
		name string
		args args
		want UserUsecase
	}{
		{
			name: "success",
			args: args{
				r: &repository_mock.MockUserRepository{},
			},
			want: UserUsecase{
				r: &repository_mock.MockUserRepository{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserUsecase(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserUsecase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserUsecase_Signin(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRepo := repository_mock.NewMockUserRepository(ctrl)
	hashedPassword, _ := hash.GenerateHashPassword("password")

	type fields struct {
		r repository.UserRepository
	}
	type args struct {
		ctx   context.Context
		param model.UserSigninParam
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.User
		wantErr bool
		mockErr error
	}{
		{
			name: "success",
			fields: fields{
				r: mockRepo,
			},
			args: args{
				ctx: context.Background(),
				param: model.UserSigninParam{
					Name:     "testuser",
					Password: "password",
				},
			},
			want: &model.User{
				ID:       1,
				Name:     "testuser",
				Password: hashedPassword,
			},
			wantErr: false,
			mockErr: nil,
		},
		{
			name: "repository error",
			fields: fields{
				r: mockRepo,
			},
			args: args{
				ctx: context.Background(),
				param: model.UserSigninParam{
					Name:     "testuser",
					Password: "wrongpassword",
				},
			},
			want:    nil,
			wantErr: true,
			mockErr: errors.New("user not found"),
		},
		{
			name: "invalid password",
			fields: fields{
				r: mockRepo,
			},
			args: args{
				ctx: context.Background(),
				param: model.UserSigninParam{
					Name:     "testuser",
					Password: "wrongpassword",
				},
			},
			want:    nil,
			wantErr: true,
			mockErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.name == "repository error" {
				mockRepo.EXPECT().GetByName(tt.args.ctx, tt.args.param.Name).Return(nil, tt.mockErr)
			} else {
				mockRepo.EXPECT().GetByName(tt.args.ctx, tt.args.param.Name).Return(&model.User{
					ID:       1,
					Name:     "testuser",
					Password: hashedPassword,
				}, nil)
			}

			u := &UserUsecase{
				r: tt.fields.r,
			}
			got, err := u.Signin(tt.args.ctx, tt.args.param)
			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestUserUsecase_GetByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRepo := repository_mock.NewMockUserRepository(ctrl)

	type fields struct {
		r repository.UserRepository
	}
	type args struct {
		ctx context.Context
		id  int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.User
		wantErr bool
		mockErr error
	}{
		{
			name: "success",
			fields: fields{
				r: mockRepo,
			},
			args: args{
				ctx: context.Background(),
				id:  1,
			},
			want: &model.User{
				ID:       1,
				Name:     "testuser",
				Password: "password",
			},
			wantErr: false,
			mockErr: nil,
		},
		{
			name: "repository error",
			fields: fields{
				r: mockRepo,
			},
			args: args{
				ctx: context.Background(),
				id:  2,
			},
			want:    nil,
			wantErr: true,
			mockErr: errors.New("user not found"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo.EXPECT().GetByID(tt.args.ctx, tt.args.id).Return(tt.want, tt.mockErr)

			u := &UserUsecase{
				r: tt.fields.r,
			}
			got, err := u.GetByID(tt.args.ctx, tt.args.id)
			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestUserUsecase_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRepo := repository_mock.NewMockUserRepository(ctrl)
	hashedPassword, _ := hash.GenerateHashPassword("password")

	type fields struct {
		r repository.UserRepository
	}
	type args struct {
		ctx   context.Context
		param model.CreateUserParam
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.User
		wantErr bool
		mockErr error
	}{
		{
			name: "success",
			fields: fields{
				r: mockRepo,
			},
			args: args{
				ctx: context.Background(),
				param: model.CreateUserParam{
					Name:     "testuser",
					Password: "password",
				},
			},
			want: &model.User{
				ID:       1,
				Name:     "testuser",
				Password: hashedPassword,
			},
			wantErr: false,
			mockErr: nil,
		},
		{
			name: "repository error",
			fields: fields{
				r: mockRepo,
			},
			args: args{
				ctx: context.Background(),
				param: model.CreateUserParam{
					Name:     "testuser",
					Password: "password",
				},
			},
			want:    nil,
			wantErr: true,
			mockErr: errors.New("create user failed"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo.EXPECT().Create(tt.args.ctx, gomock.Any()).Return(tt.want, tt.mockErr)

			u := &UserUsecase{
				r: tt.fields.r,
			}
			got, err := u.Create(tt.args.ctx, tt.args.param)
			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, tt.want, got)
		})
	}
}
