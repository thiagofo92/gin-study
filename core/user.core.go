package core

import (
	inputapp "thiagofo92/study-api-gin/app/input_app"
	"thiagofo92/study-api-gin/app/outputapp"
)

type UserCore interface {
	Create(inputapp.UserInput) (outputapp.UserOutPut, error)
	FindById(id string) (outputapp.UserOutPut, error)
	Update(id string, input inputapp.UserInput) error
}
