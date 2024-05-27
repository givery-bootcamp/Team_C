package main

import (
	"fmt"
	"myapp/internal/config"
	"myapp/internal/interface/api/router"
)

func main() {
	router := router.CreateRouter()
	router.Run(fmt.Sprintf("%s:%d", config.HostName, config.Port))
}
