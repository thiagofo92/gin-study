package controller

import (
	"encoding/json"
	"log/slog"
	"net/http"
	inputapp "thiagofo92/study-api-gin/internal/app/input_app"
	"thiagofo92/study-api-gin/internal/core"

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

// @Tags user
// @Accept		json
// @Success	201	{object}	outputapp.UserOutPut
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router		/user [post]
func (c *controller) Create(ctx *gin.Context) {
	var input inputapp.UserInput

	err := json.NewDecoder(ctx.Request.Body).Decode(&input)

	if err != nil {
		slog.Warn("error to unmarshal data %v", err)
		ctx.JSON(http.StatusBadRequest, "Bad Request")
	}

	resul, err := c.rep.Create(input)

	if err != nil {
		slog.Warn("Error to create a user %v", err)
		ctx.JSON(http.StatusInternalServerError, "Internal Server error")
	}

	ctx.JSON(http.StatusCreated, resul)
}

// @Tags user
// @Accept		json
// @Success	204
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Param id path string true "User Id"
// @Router		/user [put]
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

// @Tags user
// @Success	200	{object}	outputapp.UserOutPut
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Param id path string true "User ID"
// @Router		/user [get]
func (c *controller) FindById(ctx *gin.Context) {
	id := ctx.Param("id")

	output, err := c.rep.FindById(id)

	if err != nil {
		slog.Warn("error to unmarshal data %v", err)
		return
	}

	ctx.JSON(http.StatusOK, output)
}
