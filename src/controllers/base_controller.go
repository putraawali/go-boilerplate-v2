package controllers

import "github.com/sarulabs/di"

type Controllers struct {
	User UserController
}

func NewController(di di.Container) *Controllers {
	return &Controllers{
		User: NewUserController(di),
	}
}
