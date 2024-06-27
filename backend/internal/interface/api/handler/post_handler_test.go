package handler

import (
	"myapp/internal/application/usecase"
	"myapp/internal/domain/model"
	"myapp/internal/domain/repository/repository_mock"
	"myapp/internal/exception"
	"myapp/internal/interface/api/middleware"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
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
			u := NewPostHandler(tt.args.u)
			assert.Equal(t, tt.want.u, u.u)
		})
	}
}

func TestPostHandler_GetAll(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockRepo := repository_mock.NewMockPostRepository(ctrl)
	mockUsecase := usecase.NewPostUsecase(mockRepo)
	handler := NewPostHandler(mockUsecase)

	tests := []struct {
		name            string
		limit           string
		offset          string
		mockReturnPosts []*model.Post
		mockError       error
		expectedStatus  int
		expectedBody    string
	}{
		{
			name:   "success",
			limit:  "10",
			offset: "0",
			mockReturnPosts: []*model.Post{
				{ID: 1, Title: "Test Post", Body: "", User: model.User{ID: 1, Name: "User1"}},
			}, mockError: nil,
			expectedStatus: http.StatusOK,

			expectedBody: `[{"id":1,"title":"Test Post","body":"","created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z","user":{"id":1,"name":"User1"}}]`,
		},
		{
			name:            "invalid limit",
			limit:           "invalid",
			offset:          "0",
			mockReturnPosts: nil,
			mockError:       nil,
			expectedStatus:  http.StatusBadRequest,
			expectedBody:    `{"code":0,"message":"リクエストが不正です"}`,
		},
		{
			name:            "invalid offset",
			limit:           "10",
			offset:          "invalid",
			mockReturnPosts: nil,
			mockError:       nil,
			expectedStatus:  http.StatusBadRequest,
			expectedBody:    `{"code":0,"message":"リクエストが不正です"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mockReturnPosts != nil || tt.mockError != nil {
				mockRepo.EXPECT().GetAll(gomock.Any(), gomock.Any(), gomock.Any()).Return(tt.mockReturnPosts, tt.mockError)
			}

			gin.SetMode(gin.TestMode)
			r := gin.Default()
			r.Use(middleware.HandleError())
			r.GET("/api/posts", handler.GetAll)

			req, _ := http.NewRequest("GET", "/api/posts?limit="+tt.limit+"&offset="+tt.offset, nil)
			w := httptest.NewRecorder()

			r.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)
			if tt.expectedStatus == http.StatusBadRequest {
				assert.Contains(t, w.Body.String(), "リクエストが不正です")
			} else {
				assert.JSONEq(t, tt.expectedBody, w.Body.String())
			}
		})
	}
}

func TestPostHandler_GetByID(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockRepo := repository_mock.NewMockPostRepository(ctrl)
	mockUsecase := usecase.NewPostUsecase(mockRepo)
	handler := NewPostHandler(mockUsecase)

	tests := []struct {
		name           string
		postID         string
		mockReturnPost *model.Post
		mockError      error
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "success",
			postID:         "1",
			mockReturnPost: &model.Post{ID: 1, Title: "Test Post", Body: "", User: model.User{ID: 1, Name: "User1"}},
			mockError:      nil,
			expectedStatus: http.StatusOK,
			expectedBody:   `{"id":1,"title":"Test Post","body":"","created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z","user":{"id":1,"name":"User1"}}`,
		},
		{
			name:           "post not found",
			postID:         "1",
			mockReturnPost: nil,
			mockError:      exception.RecordNotFoundError,
			expectedStatus: http.StatusNotFound,
			expectedBody:   `{"code":0,"message":"レコードが見つかりませんでした"}`,
		},
		{
			name:           "invalid id",
			postID:         "invalid",
			mockReturnPost: nil,
			mockError:      nil,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"code":0,"message":"リクエストが不正です"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mockReturnPost != nil || tt.mockError != nil {
				mockRepo.EXPECT().GetByID(gomock.Any(), 1).Return(tt.mockReturnPost, tt.mockError)
			}

			gin.SetMode(gin.TestMode)
			r := gin.Default()
			r.Use(middleware.HandleError())
			r.GET("/api/posts/:id", handler.GetByID)

			req, _ := http.NewRequest("GET", "/api/posts/"+tt.postID, nil)
			w := httptest.NewRecorder()

			r.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)
			assert.JSONEq(t, tt.expectedBody, w.Body.String())
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
