package application

import (
	"log"

	"alexandre-gerault.fr/gochat-server/internal/messaging/domain"
	"github.com/google/uuid"
)

type SendMessageDto struct {
	author_id string
	room_id   string
	content   string
}

type SendMessagePresenter interface {
	Presents()
	MessageEmpty()
	TooLongMessage()
	AuthorNotFound()
}

type MessageRepository interface {
	Save(message domain.Message)
}

type AuthorRepository interface {
	GetById(id uuid.UUID) (domain.Author, bool)
}

type UuidProvider interface{
	Generate() uuid.UUID
}

func SendMessageHandler(
	userRepository AuthorRepository,
	messageRepository MessageRepository,
	uuidProvider UuidProvider,
) func(dto SendMessageDto, presenter SendMessagePresenter) {
	return func(dto SendMessageDto, presenter SendMessagePresenter) {
		if len(dto.content) == 0 {
			presenter.MessageEmpty()
			return
		}

		if len(dto.content) > 2000 {
			presenter.TooLongMessage()
			return
		}

		author_id, author_err := uuid.Parse(dto.author_id)
		room_id, room_err := uuid.Parse(dto.room_id)

		if author_err != nil || room_err != nil {
			log.Fatal("Invalid UUID format")
		}

		if _, found := userRepository.GetById(author_id); !found {
			presenter.AuthorNotFound()
			return
		}

		message_id := uuidProvider.Generate()
		message := domain.NewMessage(message_id, room_id, author_id, dto.content)
		messageRepository.Save(message)

		presenter.Presents()
	}
}
