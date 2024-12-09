package bootstrap

import (
	"github.com/iacopoghilardi/mydget-backend/internals/db"
	"github.com/iacopoghilardi/mydget-backend/internals/repositories"
)

type Repositories struct {
	UserRepository *repositories.UserRepository
}

func SetupRepositories() *Repositories {
	return &Repositories{
		UserRepository: repositories.NewUserRepository(db.GetDB()),
	}
}
