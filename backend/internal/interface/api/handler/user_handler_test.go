package handler

import (
	"bytes"
	"encoding/json"
	"myapp/internal/application/usecase"
	"myapp/internal/domain/model"
	"myapp/internal/domain/repository/repository_mock"
	"myapp/internal/exception"
	"myapp/internal/interface/api/middleware"
	"myapp/internal/pkg/hash"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
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
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository_mock.NewMockUserRepository(ctrl)
	mockUsecase := usecase.NewUserUsecase(mockRepo)
	handler := NewUserHandler(mockUsecase)
	hashedPassword, _ := hash.GenerateHashPassword("password")

	tests := []struct {
		name           string
		body           model.UserSigninParam
		mockReturnUser *model.User
		mockError      error
		expectedStatus int
		expectedBody   string
	}{
		{
			name: "success",
			body: model.UserSigninParam{Name: "testuser", Password: "password"},
			mockReturnUser: &model.User{
				ID:       1,
				Name:     "testuser",
				Password: hashedPassword,
			},
			mockError:      nil,
			expectedStatus: http.StatusOK,
			expectedBody:   `{"id":1,"name":"testuser"}`,
		},
		{
			name:           "internal server error",
			body:           model.UserSigninParam{Name: "testuser", Password: "password"},
			mockReturnUser: nil,
			mockError:      exception.ServerError,
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   `{"code":0,"message":"エラーが発生しました"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo.EXPECT().GetByName(gomock.Any(), tt.body.Name).Return(tt.mockReturnUser, tt.mockError)

			gin.SetMode(gin.TestMode)
			r := gin.Default()
			r.Use(middleware.HandleError())
			r.POST("/signin", handler.Signin)

			body, _ := json.Marshal(tt.body)
			req, _ := http.NewRequest("POST", "/signin", bytes.NewBuffer(body))
			w := httptest.NewRecorder()

			r.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)
			assert.JSONEq(t, tt.expectedBody, w.Body.String())
		})
	}
}

func TestUserHandler_Signout(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository_mock.NewMockUserRepository(ctrl)
	mockUsecase := usecase.NewUserUsecase(mockRepo)
	handler := NewUserHandler(mockUsecase)

	tests := []struct {
		name           string
		mockError      error
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "successful signout",
			mockError:      nil,
			expectedStatus: http.StatusOK,
			expectedBody:   `{}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)
			r := gin.Default()
			r.Use(middleware.HandleError())
			r.POST("/signout", handler.Signout)

			req, _ := http.NewRequest("POST", "/signout", nil)
			w := httptest.NewRecorder()

			r.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)
			assert.JSONEq(t, tt.expectedBody, w.Body.String())
		})
	}
}

func TestUserHandler_GetByIDFromContext(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository_mock.NewMockUserRepository(ctrl)
	mockUsecase := usecase.NewUserUsecase(mockRepo)
	handler := NewUserHandler(mockUsecase)
	// mockMiddleware := middleware_mock.NewMockAuthMiddleware(ctrl)

	tests := []struct {
		name           string
		userID         string
		mockReturnUser *model.User
		mockError      error
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "user found",
			userID:         "1",
			mockReturnUser: &model.User{ID: 1, Name: "testuser", Password: "password"},
			mockError:      nil,
			expectedStatus: http.StatusOK,
			expectedBody:   `{"id":1,"name":"testuser"}`,
		},
		{
			name:           "user not found",
			userID:         "999",
			mockReturnUser: nil,
			mockError:      exception.RecordNotFoundError,
			expectedStatus: http.StatusNotFound,
			expectedBody:   `{"code":0,"message":"レコードが見つかりませんでした"}`,
		},
		{
			name:           "internal server error",
			userID:         "1",
			mockReturnUser: nil,
			mockError:      exception.ServerError,
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   `{"code":0,"message":"エラーが発生しました"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo.EXPECT().GetByID(gomock.Any(), 1).Return(tt.mockReturnUser, tt.mockError)

			// TODO : ここを直す
			originalGetUserIDFromContext := middleware.GetUserIDFromContext
			middleware.GetUserIDFromContext = func(ctx *gin.Context) (int, error) {
				return 1, nil
			}
			defer func() {
				middleware.GetUserIDFromContext = originalGetUserIDFromContext
			}()

			gin.SetMode(gin.TestMode)
			r := gin.Default()
			r.Use(middleware.HandleError())
			r.GET("/user/:id", handler.GetByIDFromContext)

			req, _ := http.NewRequest("GET", "/user/1", nil)
			w := httptest.NewRecorder()

			r.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)
			assert.JSONEq(t, tt.expectedBody, w.Body.String())
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
