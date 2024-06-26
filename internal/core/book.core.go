package core

import (
	inputapp "thiagofo92/study-api-gin/internal/app/input_app"
	"thiagofo92/study-api-gin/internal/app/outputapp"
)

type BookCore interface {
	Add(input inputapp.BookInput) (outputapp.BookOutput, error)
	Update(id string, input inputapp.BookInput) error
	FindById(id string) (outputapp.BookOutput, error)
	Delete(id string) (int64, error)
	UpdateRent(idStr string, count int) error
}
