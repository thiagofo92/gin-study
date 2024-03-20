package controller

import (
	"log/slog"
	"net/http"
	inputapp "thiagofo92/study-api-gin/app/input_app"
	repmongo "thiagofo92/study-api-gin/infra/repository/rep_mongo"

	"github.com/gin-gonic/gin"
)

type BookController struct {
	rp *repmongo.BooksRepository
}

func NewBookController(rp *repmongo.BooksRepository) *BookController {
	return &BookController{
		rp: rp,
	}
}

func (b *BookController) Add(ctx *gin.Context) {
	var input inputapp.BookInput

	err := ctx.BindJSON(&input)

	if err != nil {
		slog.Warn("error to bind json %v", err)
		ctx.JSON(http.StatusBadRequest, "invalid input")
		return
	}

	out, err := b.rp.Add(input)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "internal server error")
		return
	}

	ctx.JSON(http.StatusCreated, out)
}

func (b *BookController) Update(ctx *gin.Context) {
	var input inputapp.BookInput

	err := ctx.BindJSON(&input)
	id := ctx.Param("id")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "invalid input")
		return
	}

	err = b.rp.Update(id, input)

	if err != nil {
		slog.Error("error to update: %v", err)
		ctx.JSON(http.StatusInternalServerError, "internal server error")
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{"message": "success"})
}

func (b *BookController) FindById(ctx *gin.Context) {
	id := ctx.Param("id")

	output, err := b.rp.FindById(id)

	if err != nil {
		slog.Error("error to update: %v", err)
		ctx.JSON(http.StatusInternalServerError, "internal server error")
		return
	}

	ctx.JSON(http.StatusNoContent, output)
}

func (b *BookController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")

	output, err := b.rp.Delete(id)

	if err != nil {
		slog.Error("error to update: %v", err)
		ctx.JSON(http.StatusInternalServerError, "internal server error")
		return
	}

	ctx.JSON(http.StatusNoContent, output)
}
