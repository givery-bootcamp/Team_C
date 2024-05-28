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
	res, err := h.u.GetAll(ctx)
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

type CreatePostRequest struct {
	Title string `json:"title" binding:"required"`
	Body  string `json:"body" binding:"required"`
}

func (h *PostHandler) Create(ctx *gin.Context) {
	var req CreatePostRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Error(exception.InvalidRequestError)
		return
	}

	newPost, err := h.u.Create(ctx, req.Title, req.Body)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusCreated, newPost)
}
