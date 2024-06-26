package handler

import (
	"errors"
	"fmt"
	"myapp/internal/application/usecase"
	"myapp/internal/exception"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type HelloWorldHandler struct {
	u usecase.HelloWorldUsecase
}

func NewHelloWorldHandler(u usecase.HelloWorldUsecase) HelloWorldHandler {
	return HelloWorldHandler{
		u: u,
	}
}

// HelloWorld godoc
//
//	@Summary	hello world
//	@Schemes
//	@Description	hello world
//	@Tags			helloWorld
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	model.HelloWorld
//	@Router			/hello [get]
func (h *HelloWorldHandler) HelloWorld(ctx *gin.Context) {
	lang := ctx.DefaultQuery("lang", "ja")
	if err := validateHelloWorldParameters(lang); err != nil {
		ctx.Error(exception.InvalidRequestError)
		return
	}
	result, err := h.u.Execute(ctx, lang)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.Error(exception.RecordNotFoundError)
		} else {
			ctx.Error(exception.ServerError)
		}
		return
	}
	ctx.JSON(200, result)
}

func validateHelloWorldParameters(lang string) error {
	if len(lang) != 2 {
		return errors.New(fmt.Sprintf("Invalid lang parameter: %s", lang))
	}
	return nil
}
