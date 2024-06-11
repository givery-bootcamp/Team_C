package main

import (
	"fmt"
	"myapp/internal/config"
	"myapp/internal/infrastructure/persistence/datastore/driver"
	"myapp/internal/interface/api/middleware"
	"myapp/internal/interface/api/router"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//	@title			掲示板アプリ
//	@version		バージョン(1.0)
//	@description	3班の掲示板アプリのAPI仕様書

//	@host		localhost:9000
//	@BasePath	/api/
func main() {
	var migrator middleware.DBMigrator = driver.MustNewMySQLMigrator("file://migrate")
	if err := migrator.Migrate(); err != nil {
		panic(err)
	}

	router := router.CreateRouter()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.Run(fmt.Sprintf("%s:%d", config.HostName, config.Port))
}
