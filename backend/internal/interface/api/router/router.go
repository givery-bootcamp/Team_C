package router

import (
	"myapp/internal/applocation/usecase"
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

	hu := usecase.NewHelloWorldUsecase(hr)

	hh := handler.NewHelloWorldHandler(hu)

	app.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "It works")
	})
	app.GET("/hello", hh.HelloWorld)

	return app
}
