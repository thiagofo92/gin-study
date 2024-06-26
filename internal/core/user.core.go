package core

import (
	inputapp "thiagofo92/study-api-gin/internal/app/input_app"
	"thiagofo92/study-api-gin/internal/app/outputapp"
)

type UserCore interface {
	Create(inputapp.UserInput) (outputapp.UserOutPut, error)
	FindById(id string) (outputapp.UserOutPut, error)
	Update(id string, input inputapp.UserInput) error
	RentBook(id string, bookId string) error
	ReturnBook(id string, bookId string) error
}
