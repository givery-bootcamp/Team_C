package usecase

import (
	"context"
	"errors"
	"myapp/internal/domain/model"
	"myapp/internal/domain/repository"
	"myapp/internal/domain/repository/repository_mock"
	"reflect"
	"testing"

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
			if got := NewHelloWorldUsecase(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHelloWorldUsecase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHelloWorldUsecase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository_mock.NewMockHelloWorldRepository(ctrl)

	type fields struct {
		r repository.HelloWorldRepository
	}
	type args struct {
		ctx  context.Context
		lang string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.HelloWorld
		wantErr bool
		setup   func()
	}{
		{
			name: "success",
			fields: fields{
				r: mockRepo,
			},
			args: args{
				ctx:  context.Background(),
				lang: "en",
			},
			want: &model.HelloWorld{
				Message: "Hello, World!",
				Lang:    "en",
			},
			wantErr: false,
			setup: func() {
				mockRepo.EXPECT().Get(context.Background(), "en").Return(&model.HelloWorld{
					Message: "Hello, World!",
					Lang:    "en",
				}, nil)
			},
		},
		{
			name: "fail/Repositoryがエラーを返した時にエラーを返す",
			fields: fields{
				r: mockRepo,
			},
			args: args{
				ctx:  context.Background(),
				lang: "en",
			},
			want:    nil,
			wantErr: true,
			setup: func() {
				mockRepo.EXPECT().Get(context.Background(), "en").Return(nil, errors.New("error"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				tt.setup()
			}
			u := &HelloWorldUsecase{
				r: tt.fields.r,
			}
			got, err := u.Execute(tt.args.ctx, tt.args.lang)
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
