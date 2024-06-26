package handler

import (
	"myapp/internal/application/usecase"
	"myapp/internal/domain/repository/repository_mock"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
	"go.uber.org/mock/gomock"
)

func TestNewPostHandler(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockRepo := repository_mock.NewMockPostRepository(ctrl)
	mockUsecase := usecase.NewPostUsecase(mockRepo)

	type args struct {
		u usecase.PostUsecase
	}
	tests := []struct {
		name string
		args args
		want PostHandler
	}{
		{
			name: "success",
			args: args{
				u: mockUsecase,
			},
			want: PostHandler{
				u: mockUsecase,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPostHandler(tt.args.u); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPostHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPostHandler_GetAll(t *testing.T) {
	type fields struct {
		u usecase.PostUsecase
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
			h := &PostHandler{
				u: tt.fields.u,
			}
			h.GetAll(tt.args.ctx)
		})
	}
}

func TestPostHandler_GetByID(t *testing.T) {
	type fields struct {
		u usecase.PostUsecase
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
			h := &PostHandler{
				u: tt.fields.u,
			}
			h.GetByID(tt.args.ctx)
		})
	}
}

func TestPostHandler_Create(t *testing.T) {
	type fields struct {
		u usecase.PostUsecase
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
			h := &PostHandler{
				u: tt.fields.u,
			}
			h.Create(tt.args.ctx)
		})
	}
}

func TestPostHandler_Update(t *testing.T) {
	type fields struct {
		u usecase.PostUsecase
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
			h := &PostHandler{
				u: tt.fields.u,
			}
			h.Update(tt.args.ctx)
		})
	}
}

func TestPostHandler_Delete(t *testing.T) {
	type fields struct {
		u usecase.PostUsecase
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
			h := &PostHandler{
				u: tt.fields.u,
			}
			h.Delete(tt.args.ctx)
		})
	}
}
