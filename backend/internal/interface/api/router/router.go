package router

import (
	"myapp/internal/application/usecase"
	"myapp/internal/infrastructure/persistence/datastore"
	"myapp/internal/infrastructure/persistence/datastore/driver"
	"myapp/internal/interface/api/handler"
	"myapp/internal/interface/api/middleware"

	"github.com/gin-gonic/gin"
)

func CreateRouter() *gin.Engine {
	app := gin.Default()
	app.Use(middleware.HandleError())
	app.Use(middleware.Cors())

	db := driver.NewDB()

	hr := datastore.NewHelloWorldRepository(db)
	pr := datastore.NewPostRepository(db)
	ur := datastore.NewUserRepository(db)

	hu := usecase.NewHelloWorldUsecase(hr)
	pu := usecase.NewPostUsecase(pr)
	uu := usecase.NewUserUsecase(ur)

	hh := handler.NewHelloWorldHandler(hu)
	ph := handler.NewPostHandler(pu)
	uh := handler.NewUserHandler(uu)

	app.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "It works")
	})
	app.GET("/hello", hh.HelloWorld)

	apiRoute := app.Group("/api")

	apiRoute.GET("/posts", ph.GetAll)
	apiRoute.GET("/posts/:id", ph.GetByID)

	apiRoute.POST("/signin", uh.Signin)

	return app
}
