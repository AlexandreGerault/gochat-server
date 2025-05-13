package infrastructure

import (
	"alexandre-gerault.fr/gochat-server/internal/messaging/domain"
	"database/sql"
)

type SqlMessageRepository struct {
	Database *sql.DB
}

func (sql_message_repository *SqlMessageRepository) Save(message domain.Message) {
}
