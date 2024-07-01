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
//	@Router			/api/signin [post]
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
//	@Router			/api/signout [post]
func (h *UserHandler) Signout(ctx *gin.Context) {
	middleware.DeleteCookie(ctx)
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
//	@Router			/api/users  [get]
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

// Signup godoc
//
//	@Summary	Signup User
//	@Schemes
//	@Description	Create User
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			body	body		model.CreateUserParam	true	"リクエスト"
//	@Success		201		{object}	model.User				"Created"
//	@Router			/api/signup [post]
func (h *UserHandler) Signup(ctx *gin.Context) {
	body := model.CreateUserParam{}
	if ctx.ShouldBindJSON(&body) != nil {
		ctx.Error(exception.InvalidRequestError)
		return
	}

	res, err := h.u.Create(ctx, body)
	if err != nil {
		ctx.Error(err)
		return
	}

	middleware.SetJWTCookie(ctx, res.ID)
	ctx.JSON(http.StatusCreated, res)
}
