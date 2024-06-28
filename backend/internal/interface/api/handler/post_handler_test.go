package handler

import (
	"bytes"
	"encoding/json"
	"myapp/internal/application/usecase"
	"myapp/internal/config"
	"myapp/internal/domain/model"
	"myapp/internal/domain/repository/repository_mock"
	"myapp/internal/exception"
	"myapp/internal/interface/api/middleware"
	"net/http"
	"net/http/httptest"
	"regexp"
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
			},
			mockError:      nil,
			expectedStatus: http.StatusOK,
			expectedBody:   `[{"id":1,"title":"Test Post","body":"","created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z","user":{"id":1,"name":"User1"}}]`,
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
		{
			name:   "exceed max limit",
			limit:  "2000",
			offset: "0",
			mockReturnPosts: []*model.Post{
				{ID: 1, Title: "Test Post", Body: "", User: model.User{ID: 1, Name: "User1"}},
			},
			mockError:      nil,
			expectedStatus: http.StatusOK,
			expectedBody:   `[{"id":1,"title":"Test Post","body":"","created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z","user":{"id":1,"name":"User1"}}]`,
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
	ctrl := gomock.NewController(t)

	mockRepo := repository_mock.NewMockPostRepository(ctrl)
	mockUsecase := usecase.NewPostUsecase(mockRepo)
	handler := NewPostHandler(mockUsecase)

	tests := []struct {
		name           string
		body           interface{}
		mockReturnPost *model.Post
		mockError      error
		expectedStatus int
		expectedBody   string
		userID         int
		userIDError    error
	}{
		{
			name: "success",
			body: model.CreatePostParam{Title: "Test Post", Body: "Test Body"},
			mockReturnPost: &model.Post{
				ID:    1,
				Title: "Test Post",
				Body:  "Test Body",
				User: model.User{
					ID:   1,
					Name: "User1",
				},
			},
			mockError:      nil,
			expectedStatus: http.StatusCreated,
			expectedBody:   `{"id":1,"title":"Test Post","body":"Test Body","created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z","user":{"id":1,"name":"User1"}}`,
			userID:         1,
			userIDError:    nil,
		},
		{
			name:           "internal server error",
			body:           model.CreatePostParam{Title: "Test Post", Body: "Test Body"},
			mockReturnPost: nil,
			mockError:      exception.ServerError,
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   `{"code":0,"message":"エラーが発生しました"}`,
			userID:         1,
			userIDError:    nil,
		},
		{
			name:           "invalid json",
			body:           "invalid json",
			mockReturnPost: nil,
			mockError:      nil,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"code":0,"message":"リクエストが不正です"}`,
			userID:         1,
			userIDError:    nil,
		},
		{
			name:           "user ID error",
			body:           model.CreatePostParam{Title: "Test Post", Body: "Test Body"},
			mockReturnPost: nil,
			mockError:      nil,
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   `{"code":0,"message":"エラーが発生しました"}`,
			userID:         0,
			userIDError:    exception.ServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.userIDError == nil {
				if tt.name == "success" || tt.name == "internal server error" {
					mockRepo.EXPECT().Create(gomock.Any(), gomock.Any()).Return(tt.mockReturnPost, tt.mockError)
				}
			}

			gin.SetMode(gin.TestMode)
			r := gin.Default()
			r.Use(middleware.HandleError())

			r.Use(func(c *gin.Context) {
				if tt.userIDError == nil {
					c.Set(config.GinSigninUserKey, tt.userID)
				} else {
					c.Error(tt.userIDError)
				}
				c.Next()
			})

			r.POST("/api/posts", handler.Create)

			var body []byte
			var err error
			if tt.name == "invalid json" {
				body = []byte(tt.body.(string))
			} else {
				body, err = json.Marshal(tt.body)
				if err != nil {
					t.Fatalf("failed to marshal body: %v", err)
				}
			}

			req, _ := http.NewRequest("POST", "/api/posts", bytes.NewBuffer(body))
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

func TestPostHandler_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository_mock.NewMockPostRepository(ctrl)
	mockUsecase := usecase.NewPostUsecase(mockRepo)
	handler := NewPostHandler(mockUsecase)

	tests := []struct {
		name           string
		postID         string
		body           interface{}
		mockGetPost    *model.Post
		mockGetErr     error
		mockUpdateErr  error
		expectedStatus int
		expectedBody   string
		userID         int
		userIDError    error
	}{
		{
			name:   "success",
			postID: "1",
			body:   model.UpdatePostParam{Title: "Updated Post", Body: "Updated Body"},
			mockGetPost: &model.Post{
				ID:    1,
				Title: "Old Title",
				Body:  "Old Body",
				User: model.User{
					ID:   1,
					Name: "User1",
				},
			},
			mockGetErr:     nil,
			mockUpdateErr:  nil,
			expectedStatus: http.StatusOK,
			expectedBody:   `{"id":1,"title":"Updated Post","body":"Updated Body","created_at":"0001-01-01T00:00:00Z","updated_at":".*","user":{"id":1,"name":"User1"}}`,
			userID:         1,
			userIDError:    nil,
		},
		{
			name:           "internal server error",
			postID:         "1",
			body:           model.UpdatePostParam{Title: "Updated Post", Body: "Updated Body"},
			mockGetPost:    &model.Post{ID: 1, Title: "Old Title", Body: "Old Body", User: model.User{ID: 1, Name: "User1"}},
			mockGetErr:     nil,
			mockUpdateErr:  exception.ServerError,
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   `{"code":0,"message":"エラーが発生しました"}`,
			userID:         1,
			userIDError:    nil,
		},
		{
			name:           "record not found",
			postID:         "1",
			body:           model.UpdatePostParam{Title: "Updated Post", Body: "Updated Body"},
			mockGetPost:    nil,
			mockGetErr:     exception.RecordNotFoundError,
			mockUpdateErr:  nil,
			expectedStatus: http.StatusNotFound,
			expectedBody:   `{"code":0,"message":"レコードが見つかりませんでした"}`,
			userID:         1,
			userIDError:    nil,
		},
		{
			name:           "invalid json",
			postID:         "1",
			body:           "invalid json",
			mockGetPost:    nil,
			mockGetErr:     nil,
			mockUpdateErr:  nil,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"code":0,"message":"リクエストが不正です"}`,
			userID:         1,
			userIDError:    nil,
		},
		{
			name:           "user ID error",
			postID:         "1",
			mockGetPost:    nil,
			mockGetErr:     nil,
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   `{"code":0,"message":"エラーが発生しました"}`,
			userID:         0,
			userIDError:    exception.ServerError,
		},
		{
			name:           "invalid postID",
			postID:         "invalid",
			body:           model.UpdatePostParam{Title: "Updated Post", Body: "Updated Body"},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"code":0,"message":"リクエストが不正です"}`,
			userID:         1,
			userIDError:    nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.userIDError == nil {
				if tt.name == "success" || tt.name == "internal server error" || tt.name == "record not found" {
					mockRepo.EXPECT().GetByID(gomock.Any(), gomock.Eq(1)).Return(tt.mockGetPost, tt.mockGetErr)
					if tt.mockGetErr == nil {
						mockRepo.EXPECT().Update(gomock.Any(), gomock.Any()).Return(tt.mockGetPost, tt.mockUpdateErr)
					}
				}
			}

			gin.SetMode(gin.TestMode)
			r := gin.Default()
			r.Use(middleware.HandleError())

			r.Use(func(c *gin.Context) {
				if tt.userIDError == nil {
					c.Set(config.GinSigninUserKey, tt.userID)
				} else {
					c.Error(tt.userIDError)
				}
				c.Next()
			})

			r.PUT("/api/posts/:id", handler.Update)

			var body []byte
			var err error
			if tt.name == "invalid json" {
				body = []byte(tt.body.(string))
			} else {
				body, err = json.Marshal(tt.body)
				if err != nil {
					t.Fatalf("failed to marshal body: %v", err)
				}
			}

			req, _ := http.NewRequest("PUT", "/api/posts/"+tt.postID, bytes.NewBuffer(body))
			w := httptest.NewRecorder()

			r.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)
			if tt.name == "success" {
				var responseBody map[string]interface{}
				if err := json.Unmarshal(w.Body.Bytes(), &responseBody); err != nil {
					t.Fatalf("failed to unmarshal response body: %v", err)
				}

				updatedAt, ok := responseBody["updated_at"].(string)
				assert.True(t, ok, "updated_at should be a string")
				assert.Regexp(t, regexp.MustCompile(`^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}\.\d{6}.*$`), updatedAt)

				delete(responseBody, "updated_at")
				expectedBody := `{"id":1,"title":"Updated Post","body":"Updated Body","created_at":"0001-01-01T00:00:00Z","user":{"id":1,"name":"User1"}}`
				var expectedBodyMap map[string]interface{}
				if err := json.Unmarshal([]byte(expectedBody), &expectedBodyMap); err != nil {
					t.Fatalf("failed to unmarshal expected body: %v", err)
				}
				assert.Equal(t, expectedBodyMap, responseBody)
			} else {
				assert.JSONEq(t, tt.expectedBody, w.Body.String())
			}
		})
	}
}

func TestPostHandler_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockRepo := repository_mock.NewMockPostRepository(ctrl)
	mockUsecase := usecase.NewPostUsecase(mockRepo)
	handler := NewPostHandler(mockUsecase)

	tests := []struct {
		name           string
		postID         string
		mockGetPost    *model.Post
		mockGetErr     error
		mockDeleteErr  error
		expectedStatus int
		expectedBody   string
		userID         int
		userIDError    error
	}{
		{
			name:           "success",
			postID:         "1",
			mockGetPost:    &model.Post{ID: 1, User: model.User{ID: 1}},
			mockGetErr:     nil,
			mockDeleteErr:  nil,
			expectedStatus: http.StatusNoContent,
			expectedBody:   "",
			userID:         1,
			userIDError:    nil,
		},
		{
			name:           "internal server error",
			postID:         "1",
			mockGetPost:    &model.Post{ID: 1, User: model.User{ID: 1}},
			mockGetErr:     nil,
			mockDeleteErr:  exception.ServerError,
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   `{"code":0,"message":"エラーが発生しました"}`,
			userID:         1,
			userIDError:    nil,
		},
		{
			name:           "record not found",
			postID:         "1",
			mockGetPost:    nil,
			mockGetErr:     exception.RecordNotFoundError,
			mockDeleteErr:  nil,
			expectedStatus: http.StatusNotFound,
			expectedBody:   `{"code":0,"message":"レコードが見つかりませんでした"}`,
			userID:         1,
			userIDError:    nil,
		},
		{
			name:           "invalid id",
			postID:         "invalid",
			mockGetPost:    nil,
			mockGetErr:     nil,
			mockDeleteErr:  nil,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"code":0,"message":"リクエストが不正です"}`,
			userID:         1,
			userIDError:    nil,
		},
		{
			name:           "user ID error",
			postID:         "1",
			mockGetPost:    nil,
			mockGetErr:     nil,
			mockDeleteErr:  nil,
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   `{"code":0,"message":"エラーが発生しました"}`,
			userID:         0,
			userIDError:    exception.ServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.userIDError == nil {
				if tt.postID != "invalid" {
					mockRepo.EXPECT().GetByID(gomock.Any(), gomock.Eq(1)).Return(tt.mockGetPost, tt.mockGetErr)
				}
				if tt.mockGetErr == nil && tt.postID != "invalid" {
					mockRepo.EXPECT().Delete(gomock.Any(), gomock.Eq(1)).Return(tt.mockDeleteErr)
				}
			}
			gin.SetMode(gin.TestMode)
			r := gin.Default()
			r.Use(middleware.HandleError())

			r.Use(func(c *gin.Context) {
				if tt.userIDError == nil {
					c.Set(config.GinSigninUserKey, tt.userID)
				} else {
					c.Error(tt.userIDError)
				}
				c.Next()
			})
			r.DELETE("/api/posts/:id", handler.Delete)

			req, _ := http.NewRequest("DELETE", "/api/posts/"+tt.postID, nil)
			w := httptest.NewRecorder()

			r.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)
			if tt.expectedStatus != http.StatusNoContent {
				assert.JSONEq(t, tt.expectedBody, w.Body.String())
			}
		})
	}
}
