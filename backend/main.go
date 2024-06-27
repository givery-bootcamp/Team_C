package main

import (
	"context"
	"fmt"
	"myapp/internal/config"
	"myapp/internal/infrastructure/persistence/datastore/driver"
	"myapp/internal/interface/api/middleware"
	"myapp/internal/interface/api/router"
	"os"
	"os/signal"
	"time"
)

//	@title			掲示板アプリ
//	@version		バージョン(1.0)
//	@description	3班の掲示板アプリのAPI仕様書

// @host	localhost:9000
func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	var migrator middleware.DBMigrator = driver.MustNewMySQLMigrator("file://migrate")
	if err := migrator.Migrate(); err != nil {
		panic(err)
	}

	go func() {
		<-ctx.Done()
		_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		os.Exit(1)
	}()

	router := router.CreateRouter()
	router.Run(fmt.Sprintf("%s:%d", config.HostName, config.Port))
}
