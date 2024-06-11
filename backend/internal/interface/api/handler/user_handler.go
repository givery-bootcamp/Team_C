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

//	@BasePath	/api

// Signin godoc
//
//	@Summary	User signin
//	@Schemes
//	@Description	Signin
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			body	body		model.UserSigninParam	true	"リクエスト"
//	@Success		200		{object}	model.User				"OK"
//	@Router			/signin [post]
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

// Signout godoc
//
//	@Summary	user signout
//	@Schemes
//	@Description	signout
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	interface{} "OK"
//	@Router			/signout [post]
func (h *UserHandler) Signout(ctx *gin.Context) {
	err := middleware.DeleteCookie(ctx)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{})
}

// GetByIDFromContext godoc
//
//	@Summary	get login user
//	@Schemes
//	@Description	get login user
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	model.User "OK"
//	@Router			/users  [get]
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
