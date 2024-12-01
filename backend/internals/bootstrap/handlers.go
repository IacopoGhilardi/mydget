package bootstrap

import (
	"github.com/iacopoghilardi/mydget-backend/internals/handlers"
)

type Handlers struct {
	UserHandler *handlers.UserHandler
	AuthHandler *handlers.AuthHandler
}

func SetupHandlers(services *Services) *Handlers {
	return &Handlers{
		UserHandler: handlers.NewUserHandler(services.UserService),
		AuthHandler: handlers.NewAuthHandler(services.AuthService),
	}
}
