package handler

import (
	"myapp/internal/application/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PostHandler struct {
	u usecase.PostUsecase
}

func NewPostHandler(u usecase.PostUsecase) PostHandler {
	return PostHandler{
		u: u,
	}
}

func (h *PostHandler) GetAll(ctx *gin.Context) {
	result, err := h.u.GetAll(ctx)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, result)
}
