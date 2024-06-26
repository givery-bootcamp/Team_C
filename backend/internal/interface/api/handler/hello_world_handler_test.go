package handler

import (
	"fmt"
	"myapp/internal/application/usecase"
	"myapp/internal/domain/model"
	"myapp/internal/domain/repository/repository_mock"
	"myapp/internal/interface/api/middleware"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"gorm.io/gorm"
)

func TestNewHelloWorldHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository_mock.NewMockHelloWorldRepository(ctrl)
	mockUsecase := usecase.NewHelloWorldUsecase(mockRepo)

	type args struct {
		u usecase.HelloWorldUsecase
	}
	tests := []struct {
		name string
		args args
		want HelloWorldHandler
	}{
		{
			name: "success",
			args: args{
				u: mockUsecase,
			},
			want: HelloWorldHandler{
				u: mockUsecase,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := NewHelloWorldHandler(tt.args.u)
			assert.Equal(t, tt.want, u)
		})
	}
}

func TestHelloWorldHandler_HelloWorld(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockRepo := repository_mock.NewMockHelloWorldRepository(ctrl)
	mockUsecase := usecase.NewHelloWorldUsecase(mockRepo)
	handler := NewHelloWorldHandler(mockUsecase)

	t.Run("invalid lang parameter", func(t *testing.T) {
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.Use(middleware.HandleError())
		r.GET("/hello", handler.HelloWorld)

		req, _ := http.NewRequest("GET", "/hello?lang=english", nil)
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.JSONEq(t, `{"code":0, "message":"リクエストが不正です"}`, w.Body.String())
	})

	tests := []struct {
		name           string
		lang           string
		mockReturn     *model.HelloWorld
		mockError      error
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "success",
			lang:           "en",
			mockReturn:     &model.HelloWorld{Message: "Hello, World!", Lang: "en"},
			mockError:      nil,
			expectedStatus: http.StatusOK,
			expectedBody:   `{"message":"Hello, World!","lang":"en"}`,
		},
		{
			name:           "repository error",
			lang:           "en",
			mockReturn:     nil,
			mockError:      fmt.Errorf("repository error"),
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   `{"code":0,"message":"エラーが発生しました"}`,
		},
		{
			name:           "record not found",
			lang:           "en",
			mockReturn:     nil,
			mockError:      gorm.ErrRecordNotFound,
			expectedStatus: http.StatusNotFound,
			expectedBody:   `{"code":0,"message":"レコードが見つかりませんでした"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo.EXPECT().Get(gomock.Any(), tt.lang).Return(tt.mockReturn, tt.mockError)

			gin.SetMode(gin.TestMode)
			r := gin.Default()
			r.Use(middleware.HandleError())
			r.GET("/hello", handler.HelloWorld)

			req, _ := http.NewRequest("GET", "/hello?lang="+tt.lang, nil)
			w := httptest.NewRecorder()

			r.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)
			assert.JSONEq(t, tt.expectedBody, w.Body.String())
		})
	}
}

func Test_validateHelloWorldParameters(t *testing.T) {
	type args struct {
		lang string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		errMsg  string
	}{
		{
			name:    "valid lang parameter",
			args:    args{lang: "en"},
			wantErr: false,
			errMsg:  "",
		},
		{
			name:    "invalid lang parameter length",
			args:    args{lang: "english"},
			wantErr: true,
			errMsg:  "Invalid lang parameter: english",
		},
		{
			name:    "empty lang parameter",
			args:    args{lang: ""},
			wantErr: true,
			errMsg:  "Invalid lang parameter: ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateHelloWorldParameters(tt.args.lang)
			if tt.wantErr {
				assert.EqualError(t, err, tt.errMsg)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
