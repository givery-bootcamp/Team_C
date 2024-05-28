package handler

import (
	"myapp/internal/application/usecase"
	"myapp/internal/domain/model"
	"myapp/internal/exception"
	"myapp/internal/interface/api/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	u usecase.UserUsecase
}

func NewUserHandler(u usecase.UserUsecase) UserHandler {
	return UserHandler{
		u: u,
	}
}

func (h *UserHandler) Signin(ctx *gin.Context) {
	body := model.UserSigninParam{}
	if ctx.ShouldBindJSON(&body) != nil {
		ctx.Error(exception.InvalidRequestError)
		return
	}

	res, err := h.u.Signin(ctx, body)
	if err != nil {
		ctx.Error(err)
		return
	}

	middleware.SetJWTCookie(ctx, res.ID)
	ctx.JSON(http.StatusOK, res)
}

func (h *UserHandler) Signout(ctx *gin.Context) {
	err := middleware.DeleteCookie(ctx)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{})
}

func (h *UserHandler) GetByIDFromContext(ctx *gin.Context) {
	userID, err := middleware.GetUserIDFromContext(ctx)
	if err != nil {
		ctx.Error(err)
		return
	}

	res, err := h.u.GetByID(ctx, userID)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, res)
}
