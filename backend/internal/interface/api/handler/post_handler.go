package handler

import (
	"myapp/internal/application/usecase"
	"myapp/internal/domain/model"
	"myapp/internal/exception"
	"myapp/internal/interface/api/middleware"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

const (
  defaultLimit = 20
  maxLimit     = 1000
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
  if limit == 0 {
    limit = defaultLimit
  } else if limit > maxLimit {
    limit = maxLimit
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

func (h *PostHandler) Create(ctx *gin.Context) {
	var param model.CreatePostParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.Error(exception.InvalidRequestError)
		return
	}

	userId, err := middleware.GetUserIDFromContext(ctx)
	if err != nil {
		ctx.Error(err)
		return
	}

	newPost, err := h.u.Create(ctx, param.Title, param.Body, userId)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusCreated, newPost)
}

func (h *PostHandler) Update(ctx *gin.Context) {
	query := ctx.Param("id")
	postID, err := strconv.Atoi(query)
	if err != nil {
		ctx.Error(exception.InvalidRequestError)
		return
	}
  
	var param model.UpdatePostParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.Error(exception.InvalidRequestError)
		return
	}

	updatedPost, err := h.u.Update(ctx, postID, param.Title, param.Body)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusCreated, updatedPost)
}
