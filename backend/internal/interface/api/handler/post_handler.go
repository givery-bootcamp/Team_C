package handler

import (
	"myapp/internal/application/usecase"
	"myapp/internal/exception"
	"net/http"
	"strconv"

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
  limit, err := strconv.Atoi(ctx.Query("limit"))
  if err != nil {
		ctx.Error(exception.InvalidRequestError)
		return
	}
  offset, err := strconv.Atoi(ctx.Query("offset"))
  if err != nil {
		ctx.Error(exception.InvalidRequestError)
		return
	}
	res, err := h.u.GetAll(ctx, limit, offset)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func (h *PostHandler) GetByID(ctx *gin.Context) {
	query := ctx.Param("id")
	postID, err := strconv.Atoi(query)
	if err != nil {
		ctx.Error(exception.InvalidRequestError)
		return
	}

	res, err := h.u.GetByID(ctx, postID)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, res)
}
