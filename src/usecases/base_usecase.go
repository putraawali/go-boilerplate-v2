package usecases

import godi "github.com/putraawali/go-di"

type Usecases struct {
	User UserUsecase
}

func NewUsecase(di godi.Container) *Usecases {
	return &Usecases{
		User: NewUserUsecase(di),
	}
}
