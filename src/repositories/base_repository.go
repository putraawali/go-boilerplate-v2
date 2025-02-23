package repositories

import godi "github.com/putraawali/go-di"

type Repositories struct {
	User UserRepository
}

// Initiate repository layer, accept dependency injection as parameter and return *Repositories
func NewRepository(di godi.Container) *Repositories {
	return &Repositories{
		User: NewUserRepository(di),
	}
}
