package handler

import (
	"errors"
	"fmt"
	"myapp/internal/application/usecase"

	"github.com/gin-gonic/gin"
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
		ctx.Error(err)
		return
	}
	result, err := h.u.Execute(ctx, lang)
	if err != nil {
		ctx.Error(err)
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
