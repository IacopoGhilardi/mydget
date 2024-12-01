package bootstrap

import (
	"github.com/iacopoghilardi/mydget-backend/internals/db"
	"github.com/iacopoghilardi/mydget-backend/internals/repositories"
)

type Repositories struct {
	UserRepository *repositories.UserRepository
	// Altri repository qui
}

func SetupRepositories() *Repositories {
	return &Repositories{
		UserRepository: repositories.NewUserRepository(db.GetDB()),
		// Inizializza altri repository
	}
}
