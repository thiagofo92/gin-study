package web

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	_ "thiagofo92/study-api-gin/cmd/docs"
	"thiagofo92/study-api-gin/internal/web/routers"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RunServer() {
	gin.ForceConsoleColor()
	r := gin.Default()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/doc/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router := routers.NewGinRouter(r.Group("/api"))

	router.Build()
	err := r.Run(":3500")

	if err != nil {
		slog.Error("error to run server %v", err)
	}
}
