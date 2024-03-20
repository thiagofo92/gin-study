package controller

import (
	"log/slog"
	"net/http"
	inputapp "thiagofo92/study-api-gin/app/input_app"
	"thiagofo92/study-api-gin/core"

	"github.com/gin-gonic/gin"
)

type controller struct {
	rep core.UserCore
}

func NewUserController(rp core.UserCore) *controller {
	return &controller{
		rep: rp,
	}
}

func (c *controller) Create(ctx *gin.Context) {
	var input inputapp.UserInput
	err := ctx.BindJSON(&input)

	if err != nil {
		slog.Warn("error to unmarshal data %v", err)
		ctx.JSON(http.StatusBadRequest, "invalid input")
		return
	}

	resul, err := c.rep.Create(input)

	if err != nil {
		slog.Warn("error to unmarshal data %v", err)
		return
	}

	ctx.JSON(http.StatusCreated, resul)
}

func (c *controller) Update(ctx *gin.Context) {
	var input inputapp.UserInput
	err := ctx.BindJSON(&input)
	id := ctx.Param("id")
	if err != nil {
		slog.Warn("error to unmarshal data %v", err)
		ctx.JSON(http.StatusBadRequest, "invalid input")
		return
	}

	err = c.rep.Update(id, input)

	if err != nil {
		slog.Warn("error to unmarshal data %v", err)
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{"data": "success"})
}

func (c *controller) FindById(ctx *gin.Context) {
	id := ctx.Param("id")

	output, err := c.rep.FindById(id)

	if err != nil {
		slog.Warn("error to unmarshal data %v", err)
		return
	}

	ctx.JSON(http.StatusOK, output)
}
