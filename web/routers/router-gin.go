package routers

import (
	"encoding/json"
	"io"
	"log/slog"
	"thiagofo92/study-api-gin/app/controller"
	repmongo "thiagofo92/study-api-gin/infra/repository/rep_mongo"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type ginRouter struct {
	r  *gin.RouterGroup
	db *mongo.Database
}

type FnController[T interface{}] func(input T) error

type DataType[T interface{}] struct {
	Typ   string
	Value T
}

func AdapterGin[T interface{}](fn FnController[T]) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		buff, err := io.ReadAll(ctx.Request.Body)

		if err != nil {
			slog.Warn("error to read body %v", err)
		}
		var body T

		err = json.Unmarshal(buff, &body)

		if err != nil {
			slog.Warn("error to unmarshal json")
			ctx.JSON(400, err)
			return
		}

		fn(body)

		ctx.JSON(200, "Success")
	}
}

func NewGinRouter(r *gin.Engine) *ginRouter {
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

// func parseData[T interface{}](args ...T) {
// 	for _, value := range args {
// 	}
// }
