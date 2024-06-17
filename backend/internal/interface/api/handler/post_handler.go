package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"myapp/internal/application/usecase"
	"myapp/internal/domain/model"
	"myapp/internal/exception"
	"myapp/internal/interface/api/middleware"
)

const (
	defaultLimit  = 20
	maxLimit      = 1000
	defaultOffset = 0
)

type PostHandler struct {
	u usecase.PostUsecase
}

func NewPostHandler(u usecase.PostUsecase) PostHandler {
	return PostHandler{
		u: u,
	}
}

// GetAll godoc
//
//	@Summary	get posts
//	@Schemes
//	@Description	get posts
//	@Tags			post
//	@Accept			json
//	@Produce		json
//	@Param			limit	query		number	false	"Limit"
//	@Param			offset	query		number	false	"Offset"
//	@Success		200		{object}	[]model.Post
//	@Router			/api/posts [get]
func (h *PostHandler) GetAll(ctx *gin.Context) {
	limit := defaultLimit
	limitQuery := ctx.Query("limit")

	if limitQuery != "" {
		l, err := strconv.Atoi(limitQuery)
		if err != nil {
			ctx.Error(exception.InvalidRequestError)
			return
		}

		limit = l
	}

	if limit > maxLimit {
		limit = maxLimit
	}

	offset := defaultOffset
	offsetQuery := ctx.Query("offset")

	if offsetQuery != "" {
		o, err := strconv.Atoi(offsetQuery)
		if err != nil {
			ctx.Error(exception.InvalidRequestError)
			return
		}

		offset = o
	}

	res, err := h.u.GetAll(ctx, limit, offset)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, res)
}

// GetByID godoc
//
//	@Summary	get post by id
//	@Schemes
//	@Description	get post by id
//	@Tags			post
//	@Accept			json
//	@Produce		json
//	@Param			id	path		number	true	"PostID"
//	@Success		200	{object}	model.Post
//	@Router			/api/posts/{id} [get]
func (h *PostHandler) GetByID(ctx *gin.Context) {
	query := ctx.Param("id")
	postID, err := strconv.Atoi(query)
	if err != nil {
		ctx.Error(exception.InvalidRequestError)
		return
	}

	userId, err := middleware.GetUserIDFromContext(ctx)
	if err != nil {
		ctx.Error(err)
		return
	}

	res, err := h.u.GetByID(ctx, postID, userId)
	if err != nil {
		ctx.Error(exception.InvalidRequestError)
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

	userId, err := middleware.GetUserIDFromContext(ctx)
	if err != nil {
		ctx.Error(err)
		return
	}

	var param model.UpdatePostParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.Error(exception.InvalidRequestError)
		return
	}

	updatedPost, err := h.u.Update(ctx, param.Title, param.Body, postID, userId)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, updatedPost)
}
