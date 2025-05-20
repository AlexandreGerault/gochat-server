package infrastructure

import (
	"database/sql"

	"alexandre-gerault.fr/gochat-server/internal/messaging/domain"
	"github.com/google/uuid"
)

type SqlMessageRepository struct {
	Database *sql.DB
}

func (sql_message_repository SqlMessageRepository) Save(message domain.Message) (uuid.UUID, error) {
	query := `
		INSERT INTO messages (id, room_id, author_id, content)
		VALUES ($1, $2, $3, $4)
	`
	_, err := sql_message_repository.Database.Exec(
		query,
		message.Id,
		message.Room_Id,
		message.Author_Id,
		message.Content,
	)

	if err != nil {
		return uuid.Nil, err
	}

	return message.Id, nil
}
