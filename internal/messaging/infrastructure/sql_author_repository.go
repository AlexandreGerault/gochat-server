package infrastructure

import (
	"database/sql"

	"alexandre-gerault.fr/gochat-server/internal/messaging/domain"
	"github.com/google/uuid"
)

type SqlAuthorRepository struct {
	Database *sql.DB
}

func (sql_author_repository *SqlAuthorRepository) Exist(id uuid.UUID) bool {
	row := sql_author_repository.Database.QueryRow("SELECT id FROM users WHERE id = $1", id.String())

	var author domain.Author

	if err := row.Scan(&author.Uuid); err != nil {
		return false
	}

	return true
}
