package shared_infrastructure

import (
	"database/sql"

	"alexandre-gerault.fr/gochat-server/internal/messaging/application"
	"alexandre-gerault.fr/gochat-server/internal/messaging/domain"
	"alexandre-gerault.fr/gochat-server/internal/messaging/infrastructure"
)

type Application struct {
	Database     *sql.DB
	Dependencies Dependencies
}

type Dependencies struct {
	AuthorRepository  domain.AuthorRepository
	MessageRepository domain.MessageRepository
	UuidProvider      application.UuidProvider
}

func (app *Application) Register() *Application {
	app.Database = CreateDatabase()
	app.Dependencies.MessageRepository = &infrastructure.SqlMessageRepository{Database: app.Database}
	app.Dependencies.AuthorRepository = &infrastructure.SqlAuthorRepository{Database: app.Database}
	app.Dependencies.UuidProvider = &UuidGenerator{}

	return app
}
