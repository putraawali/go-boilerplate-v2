package controllers

import godi "github.com/putraawali/go-di"

type Controllers struct {
	User UserController
}

func NewController(di godi.Container) *Controllers {
	return &Controllers{
		User: NewUserController(di),
	}
}
