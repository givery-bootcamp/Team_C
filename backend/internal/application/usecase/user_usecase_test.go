package usecase

import (
	"context"
	"myapp/internal/domain/model"
	"myapp/internal/domain/repository"
	"myapp/internal/domain/repository/repository_mock"
	"reflect"
	"testing"
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
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserUsecase{
				r: tt.fields.r,
			}
			got, err := u.Signin(tt.args.ctx, tt.args.param)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserUsecase.Signin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserUsecase.Signin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserUsecase_GetByID(t *testing.T) {
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
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserUsecase{
				r: tt.fields.r,
			}
			got, err := u.GetByID(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserUsecase.GetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserUsecase.GetByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserUsecase_Create(t *testing.T) {
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
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserUsecase{
				r: tt.fields.r,
			}
			got, err := u.Create(tt.args.ctx, tt.args.param)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserUsecase.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserUsecase.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}
