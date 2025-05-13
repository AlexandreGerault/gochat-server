package infrastructure

import (
	"database/sql"
	"fmt"

	"alexandre-gerault.fr/gochat-server/internal/messaging/domain"
	"github.com/google/uuid"
)

type SqlAuthorRepository struct {
	Database *sql.DB
}

func (sql_author_repository *SqlAuthorRepository) GetById(id uuid.UUID) (domain.Author, error) {
	row := sql_author_repository.Database.QueryRow("SELECT uuid FROM authors WHERE uuid = $1", id.String())

	var author domain.Author
		
	if err := row.Scan(&author.Uuid); err != nil {
		if err == sql.ErrNoRows {
			return author, fmt.Errorf("Cannot find author %s", id.String())
		}

		return author, fmt.Errorf("Error finding author (%s): %s", id, err)
	}

	return author, nil
}

