package routers

import (
	"thiagofo92/study-api-gin/internal/app/controller"
	repmongo "thiagofo92/study-api-gin/internal/infra/repository/rep_mongo"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type ginRouter struct {
	r  *gin.RouterGroup
	db *mongo.Database
}

func NewGinRouter(r *gin.RouterGroup) *ginRouter {
	db, err := repmongo.NewConnect()

	if err != nil {
		panic(err)
	}

	return &ginRouter{
		r:  r.Group("/v1"),
		db: db,
	}
}

func (gr *ginRouter) user() {
	rep := repmongo.NewUserRepository(gr.db)
	controller := controller.NewUserController(rep)

	gr.r.POST("/user", controller.Create)
	gr.r.GET("/user/:id", controller.FindById)
	gr.r.PUT("/user/:id", controller.Update)
}

func (gr *ginRouter) book() {
	rep := repmongo.NewBooksRepository(gr.db)
	controller := controller.NewBookController(rep)

	gr.r.POST("/book", controller.Add)
	gr.r.PUT("/book/:id", controller.Update)
	gr.r.GET("/book/:id", controller.FindById)
	gr.r.DELETE("/book/:id", controller.Delete)
}

func (gr *ginRouter) Build() error {
	gr.user()
	gr.book()

	return nil
}
