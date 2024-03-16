package controller

import (
	"log/slog"
	"net/http"
	inputapp "thiagofo92/study-api-gin/app/input_app"
	repmongo "thiagofo92/study-api-gin/infra/repository/rep_mongo"

	"github.com/gin-gonic/gin"
)

type controller struct {
	rep *repmongo.UserRepository
}

func NewUserController(rp *repmongo.UserRepository) *controller {
	return &controller{
		rep: rp,
	}
}

func (c *controller) Create(ctx *gin.Context) {
	var input inputapp.UserInput
	err := ctx.BindJSON(&input)

	if err != nil {
		slog.Warn("error to unmarshal data %v", err)
		return
	}

	resul, err := c.rep.Create(input)

	if err != nil {
		slog.Warn("error to unmarshal data %v", err)
		return
	}

	ctx.JSON(http.StatusCreated, resul)
}
