package web

import (
	"log/slog"
	"thiagofo92/study-api-gin/web/routers"

	"github.com/gin-gonic/gin"
)

func RunServer() {
	gin.ForceConsoleColor()
	r := gin.Default()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	router := routers.NewGinRouter(r)

	router.Build()

	err := r.Run(":3500")

	if err != nil {
		slog.Error("error to run server %v", err)
	}
}
