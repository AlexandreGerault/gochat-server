package infrastructure

import (
	"alexandre-gerault.fr/gochat-server/internal/messaging/domain"
	"github.com/google/uuid"
)

type SqlAuthorRepository struct {}

func (sql_author_repository *SqlAuthorRepository) GetById(id uuid.UUID) (domain.Author, bool) {
	return domain.Author{}, true
}

