package main

import (
	"fmt"
	"myapp/internal/config"
	"myapp/internal/infrastructure/persistence/datastore"
	"myapp/internal/interface/api/middleware"
	"myapp/internal/interface/api/router"
)

func main() {
	var migrator middleware.DBMigrator = datastore.MustNewMySQLMigrator("file://migrate")
	if err := migrator.Migrate(); err != nil {
		panic(err)
	}

	router := router.CreateRouter()
	router.Run(fmt.Sprintf("%s:%d", config.HostName, config.Port))
}
