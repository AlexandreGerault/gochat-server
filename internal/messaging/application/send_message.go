package application

import (
	"alexandre-gerault.fr/gochat-server/internal/messaging/domain"
	"github.com/google/uuid"
)

type SendMessageDto struct {
	Author_id string
	Room_id   string
	Content   string
}

type SendMessagePresenter interface {
	Presents()
	MessageEmpty()
	TooLongMessage()
	AuthorNotFound()
	InvalidPayload()
	UnexpectedError(error string)
}

type UuidProvider interface {
	Generate() (uuid.UUID, error)
}

func SendMessageHandler(
	userRepository domain.AuthorRepository,
	messageRepository domain.MessageRepository,
	uuidProvider UuidProvider,
) func(dto SendMessageDto, presenter SendMessagePresenter) {
	return func(dto SendMessageDto, presenter SendMessagePresenter) {
		author_id, author_err := uuid.Parse(dto.Author_id)
		room_id, room_err := uuid.Parse(dto.Room_id)

		if author_err != nil || room_err != nil {
			presenter.InvalidPayload()
			return
		}

		if len(dto.Content) == 0 {
			presenter.MessageEmpty()
			return
		}

		if len(dto.Content) > 2000 {
			presenter.TooLongMessage()
			return
		}

		if _, found := userRepository.GetById(author_id); !found {
			presenter.AuthorNotFound()
			return
		}

		message_id, err := uuidProvider.Generate()

		if err != nil {
			presenter.UnexpectedError(err.Error())
		}

		message := domain.NewMessage(message_id, room_id, author_id, dto.Content)
		messageRepository.Save(message)

		presenter.Presents()
	}
}
