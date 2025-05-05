package shared_infrastructure

import (
	"alexandre-gerault.fr/gochat-server/internal/messaging/application"
	"alexandre-gerault.fr/gochat-server/internal/messaging/domain"
	"alexandre-gerault.fr/gochat-server/internal/messaging/infrastructure"
	"github.com/golobby/container/v3"
)

type Application struct {
	_container *container.Container
}

func (app *Application) Register() *Application {
	app._container.Singleton(func() domain.MessageRepository {
		return &infrastructure.SqlMessageRepository{}
	})

	app._container.Singleton(func() domain.AuthorRepository {
		return &infrastructure.SqlAuthorRepository{}
	})

	app._container.Singleton(func() application.UuidProvider {
		return &UuidGenerator{}
	})

	return app
}

func (app *Application) Container() *container.Container {
	return app._container
}
