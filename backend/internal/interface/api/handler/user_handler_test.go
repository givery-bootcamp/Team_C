package handler

import (
	"myapp/internal/application/usecase"
	"myapp/internal/domain/repository/repository_mock"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
	"go.uber.org/mock/gomock"
)

func TestNewUserHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRepo := repository_mock.NewMockUserRepository(ctrl)
	mockUsecase := usecase.NewUserUsecase(mockRepo)

	type args struct {
		u usecase.UserUsecase
	}
	tests := []struct {
		name string
		args args
		want UserHandler
	}{
		{
			name: "success",
			args: args{
				u: mockUsecase,
			},
			want: UserHandler{
				u: mockUsecase,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserHandler(tt.args.u); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserHandler_Signin(t *testing.T) {
	type fields struct {
		u usecase.UserUsecase
	}
	type args struct {
		ctx *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &UserHandler{
				u: tt.fields.u,
			}
			h.Signin(tt.args.ctx)
		})
	}
}

func TestUserHandler_Signout(t *testing.T) {
	type fields struct {
		u usecase.UserUsecase
	}
	type args struct {
		ctx *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &UserHandler{
				u: tt.fields.u,
			}
			h.Signout(tt.args.ctx)
		})
	}
}

func TestUserHandler_GetByIDFromContext(t *testing.T) {
	type fields struct {
		u usecase.UserUsecase
	}
	type args struct {
		ctx *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &UserHandler{
				u: tt.fields.u,
			}
			h.GetByIDFromContext(tt.args.ctx)
		})
	}
}

func TestUserHandler_Signup(t *testing.T) {
	type fields struct {
		u usecase.UserUsecase
	}
	type args struct {
		ctx *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &UserHandler{
				u: tt.fields.u,
			}
			h.Signup(tt.args.ctx)
		})
	}
}
