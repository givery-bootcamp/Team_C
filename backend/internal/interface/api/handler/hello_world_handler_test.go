package handler

import (
	"myapp/internal/application/usecase"
	"myapp/internal/domain/model"
	"myapp/internal/domain/repository/repository_mock"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mockError == nil {
				mockRepo.EXPECT().Get(gomock.Any(), tt.lang).Return(tt.mockReturn, nil).AnyTimes()
			} else {
				mockRepo.EXPECT().Get(gomock.Any(), tt.lang).Return(nil, tt.mockError).AnyTimes()
			}

			gin.SetMode(gin.TestMode)
			r := gin.Default()
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
	}{
		{
			name:    "valid lang parameter",
			args:    args{lang: "en"},
			wantErr: false,
		},
		{
			name:    "invalid lang parameter length",
			args:    args{lang: "english"},
			wantErr: true,
		},
		{
			name:    "empty lang parameter",
			args:    args{lang: ""},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateHelloWorldParameters(tt.args.lang)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}
