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
	Author_Repository  domain.AuthorRepository
	Message_Repository domain.MessageRepository
	Uuid_Provider      application.UuidProvider
}

func (app *Application) Register() *Application {
	app.Database = CreateDatabase()
	app.Dependencies.Message_Repository = &infrastructure.SqlMessageRepository{Database: app.Database}
	app.Dependencies.Author_Repository = &infrastructure.SqlAuthorRepository{Database: app.Database}
	app.Dependencies.Uuid_Provider = &UuidGenerator{}

	return app
}
