package entity

import "myapp/internal/domain/model"

type HelloWorld struct {
	Lang    string
	Message string
}

func (h HelloWorld) ToModel() *model.HelloWorld {
	return &model.HelloWorld{
		Lang:    h.Lang,
		Message: h.Message,
	}
}
