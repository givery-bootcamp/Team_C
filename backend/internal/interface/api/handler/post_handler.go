package handler

import (
	"myapp/internal/application/usecase"
	"myapp/internal/exception"
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

	res, err := h.u.GetByID(ctx, postID)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, res)
}
