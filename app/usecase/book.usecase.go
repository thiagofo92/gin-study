package usecase

import "thiagofo92/study-api-gin/core"

type usecase struct {
	ruser core.UserCore
	rbook core.BookCore
}

func NewBookUseCase(ruser core.UserCore, rbook core.BookCore) *usecase {
	return &usecase{
		ruser: ruser,
		rbook: rbook,
	}
}

func (u *usecase) Rent() {

}

func (u *usecase) GiveBack() {

}
