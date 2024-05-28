package handler

import (
	"myapp/internal/application/usecase"
	"myapp/internal/domain/model"
	"myapp/internal/exception"
	"myapp/internal/interface/api/middleware"
	"myapp/internal/pkg/jwt"
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

	token, err := jwt.GenerateToken(res.ID)
	if err != nil {
		ctx.Error(err)
		return
	}
	middleware.SetCookie(ctx, token)
	ctx.JSON(http.StatusOK, res)
}
