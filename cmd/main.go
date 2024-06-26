package main

import "thiagofo92/study-api-gin/internal/web"

// @title           Swagger Example API
// @version         0.2.0
// @description     This is a sample server celler server.

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @host      localhost:3500
// @BasePath  /api/v1
func main() {
	web.RunServer()
}
