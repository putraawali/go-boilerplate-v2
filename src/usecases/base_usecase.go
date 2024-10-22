package usecases

import "github.com/sarulabs/di"

type Usecases struct {
	User UserUsecase
}

func NewUsecase(di di.Container) *Usecases {
	return &Usecases{
		User: NewUserUsecase(di),
	}
}
