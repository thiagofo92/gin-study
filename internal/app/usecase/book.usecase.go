package usecase

import (
	"thiagofo92/study-api-gin/core"
)

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

func (u *usecase) Rent(userId string, bookId string) error {
	rent := 1
	err := u.rbook.UpdateRent(bookId, rent)

	if err != nil {
		return err
	}

	err = u.ruser.RentBook(userId, bookId)

	if err != nil {
		return err
	}

	return nil
}

func (u *usecase) ReturnBook(userId string, bookId string) error {
	giveBack := -1

	err := u.rbook.UpdateRent(bookId, giveBack)

	if err != nil {
		return err
	}

	err = u.ruser.ReturnBook(userId, bookId)

	if err != nil {
		return err
	}

	return nil
}
