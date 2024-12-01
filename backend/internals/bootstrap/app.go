package bootstrap

type Application struct {
	Repositories *Repositories
	Services     *Services
	Handlers     *Handlers
}

func NewApplication() (*Application, error) {
	if err := SetupDatabase(); err != nil {
		return nil, err
	}

	repos := SetupRepositories()
	services := SetupServices(repos)
	handlers := SetupHandlers(services)

	return &Application{
		Repositories: repos,
		Services:     services,
		Handlers:     handlers,
	}, nil
}
