package main

import (
	"context"
	"fmt"
	"log"
	"myapp/internal/config"
	"myapp/internal/infrastructure/persistence/datastore/driver"
	"myapp/internal/interface/api/middleware"
	"myapp/internal/interface/api/router"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

//	@title			掲示板アプリ
//	@version		バージョン(1.0)
//	@description	3班の掲示板アプリのAPI仕様書

// @host	localhost:9000
func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	var migrator middleware.DBMigrator = driver.MustNewMySQLMigrator("file://migrate")
	if err := migrator.Migrate(); err != nil {
		log.Fatalf("Failed to migrate: %v", err)
	}

	router := router.CreateRouter()

	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", config.HostName, config.Port),
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	<-ctx.Done()
	stop()
	log.Println("Shutting down server...")

	ctxShutdown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctxShutdown); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}
	log.Println("Server exiting")
}
