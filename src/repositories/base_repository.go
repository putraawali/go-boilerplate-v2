package repositories

import "github.com/sarulabs/di"

type Repositories struct {
	User UserRepository
}

// Initiate repository layer, accept dependency injection as parameter and return *Repositories
func NewRepository(di di.Container) *Repositories {
	return &Repositories{
		User: NewUserRepository(di),
	}
}
