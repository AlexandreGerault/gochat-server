package infrastructure

import "alexandre-gerault.fr/gochat-server/internal/messaging/domain"

type SqlMessageRepository struct{}

func (sql_message_repository *SqlMessageRepository) Save(message domain.Message) {
}
