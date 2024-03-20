package web

import (
	"net/http/httptest"
	"testing"
	"thiagofo92/study-api-gin/infra/web/routers"

	"github.com/gin-gonic/gin"
)

func TestE2E(t *testing.T) {
	g := gin.Default()
	rt := routers.NewGinRouter(g)

	rt.Build()
	server := httptest.NewServer(g)

	server.Start()

	server.Close()
}
