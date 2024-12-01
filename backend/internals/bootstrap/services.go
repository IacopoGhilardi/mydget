package bootstrap

import "github.com/iacopoghilardi/mydget-backend/internals/services"

type Services struct {
	UserService *services.UserService
	AuthService *services.AuthService
}

func SetupServices(repositories *Repositories) *Services {
	return &Services{
		UserService: services.NewUserService(repositories.UserRepository),
		AuthService: services.NewAuthService(repositories.UserRepository),
	}
}
