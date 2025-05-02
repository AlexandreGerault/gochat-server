package application

import (
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

type UserRepository interface {
	GetById(id uuid.UUID) (domain.Author, bool)
}

func SendMessageHandler(userRepository UserRepository, messageRepository MessageRepository) func(dto SendMessageDto, presenter SendMessagePresenter) {
	return func(dto SendMessageDto, presenter SendMessagePresenter) {
		if len(dto.content) == 0 {
			presenter.MessageEmpty()
			return
		}

		if len(dto.content) > 2000 {
			presenter.TooLongMessage()
			return
		}

		author_id, _ := uuid.Parse(dto.author_id)

		if _, found := userRepository.GetById(author_id); found == false {
			presenter.AuthorNotFound()
			return
		}

		message := domain.NewMessage(dto.content)
		messageRepository.Save(message)

		presenter.Presents()
	}
}
