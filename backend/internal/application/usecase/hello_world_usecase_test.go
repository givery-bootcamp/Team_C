package usecase

import (
	"context"
	"errors"
	"myapp/internal/domain/model"
	"myapp/internal/domain/repository"
	"myapp/internal/domain/repository/repository_mock"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestNewHelloWorldUsecase(t *testing.T) {
	type args struct {
		r repository.HelloWorldRepository
	}
	tests := []struct {
		name string
		args args
		want HelloWorldUsecase
	}{
		{
			name: "success",
			args: args{
				r: &repository_mock.MockHelloWorldRepository{},
			},
			want: HelloWorldUsecase{
				r: &repository_mock.MockHelloWorldRepository{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := NewHelloWorldUsecase(tt.args.r)
			assert.Equal(t, u, tt.want)
		})
	}
}

func TestHelloWorldUsecase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockRepo := repository_mock.NewMockHelloWorldRepository(ctrl)

	tests := []struct {
		name    string
		lang    string
		want    *model.HelloWorld
		wantErr bool
		mockErr error
	}{
		{
			name: "success",
			lang: "en",
			want: &model.HelloWorld{
				Message: "Hello, World!",
				Lang:    "en",
			},
			wantErr: false,
			mockErr: nil,
		},
		{
			name:    "fail/Repositoryがエラーを返した時にエラーを返す",
			lang:    "en",
			want:    nil,
			wantErr: true,
			mockErr: errors.New("error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mockErr == nil {
				mockRepo.EXPECT().Get(context.Background(), tt.lang).Return(tt.want, nil)
			} else {
				mockRepo.EXPECT().Get(context.Background(), tt.lang).Return(nil, tt.mockErr)
			}

			u := HelloWorldUsecase{
				r: mockRepo,
			}
			got, err := u.Execute(context.Background(), tt.lang)
			if (err != nil) != tt.wantErr {
				t.Errorf("HelloWorldUsecase.Execute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HelloWorldUsecase.Execute() = %v, want %v", got, tt.want)
			}
		})
	}
}
