package application

import (
	"testing"

	"alexandre-gerault.fr/gochat-server/internal/messaging/domain"
	"github.com/fufuok/random"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

type SendMessageTestPresenter struct {
	response string
}

type InMemoryMessageRepository struct {
	messages []domain.Message
}

type InMemoryAuthorRepository struct {
	users []domain.Author
}

func (inMemoryUserRepository *InMemoryAuthorRepository) GetById(id uuid.UUID) (domain.Author, bool) {
	for _, author := range inMemoryUserRepository.users {
		if author.Uuid.String() == id.String() {
			return author, true
		}
	}

	return domain.Author{}, false
}

func (inMemoryMessageRepository *InMemoryMessageRepository) Save(message domain.Message) {
	inMemoryMessageRepository.messages = append(inMemoryMessageRepository.messages, message)
}

func (presenter *SendMessageTestPresenter) Presents() {
	presenter.response = "success"
}

func (presenter *SendMessageTestPresenter) MessageEmpty() {
	presenter.response = "empty"
}

func (presenter *SendMessageTestPresenter) TooLongMessage() {
	presenter.response = "too_long"
}

func (presenter *SendMessageTestPresenter) AuthorNotFound() {
	presenter.response = "author_not_found"
}

func TestItCanSendMessage(t *testing.T) {
	author_id, err := uuid.Parse("01968e00-1b4d-7a91-bb2a-c55bd56a2dac")

	dto := SendMessageDto{author_id.String(), "room_id", "Some message"}

	if err != nil {
		t.Error("Cannot generate author_id")
	}

	message_repository := InMemoryMessageRepository{}
	authors_repository := InMemoryAuthorRepository{
		[]domain.Author{domain.NewAuthor(author_id)},
	}

	presenter := SendMessageTestPresenter{}

	handler := SendMessageHandler(&authors_repository, &message_repository)

	handler(dto, &presenter)

	assert.Equal(t, "success", presenter.response)
	assert.Equal(t, 1, len(message_repository.messages))
}

func TestItCannotSendAnEmptyMessage(t *testing.T) {
	author_id, err := uuid.Parse("01968e00-1b4d-7a91-bb2a-c55bd56a2dac")

	dto := SendMessageDto{author_id.String(), "room_id", ""}

	if err != nil {
		t.Error("Cannot generate author_id")
	}

	message_repository := InMemoryMessageRepository{}
	authors_repository := InMemoryAuthorRepository{
		[]domain.Author{domain.NewAuthor(author_id)},
	}

	presenter := SendMessageTestPresenter{}

	handler := SendMessageHandler(&authors_repository, &message_repository)

	handler(dto, &presenter)

	assert.Equal(t, "empty", presenter.response)
	assert.Equal(t, 0, len(message_repository.messages))
}

func TestItCannotSendAnOversizedMessage(t *testing.T) {
	author_id, err := uuid.Parse("01968e00-1b4d-7a91-bb2a-c55bd56a2dac")

	dto := SendMessageDto{author_id.String(), "room_id", random.RandString(2001)}

	if err != nil {
		t.Error("Cannot generate author_id")
	}

	message_repository := InMemoryMessageRepository{}
	authors_repository := InMemoryAuthorRepository{
		[]domain.Author{domain.NewAuthor(author_id)},
	}

	presenter := SendMessageTestPresenter{}

	handler := SendMessageHandler(&authors_repository, &message_repository)

	handler(dto, &presenter)

	assert.Equal(t, "too_long", presenter.response)
	assert.Equal(t, 0, len(message_repository.messages))
}

func TestItCannotSendMessageIfAuthorDoesNotExist(t *testing.T) {
	dto := SendMessageDto{"user_id", "room_id", "Some legal content"}

	message_repository := InMemoryMessageRepository{}
	authors_repository := InMemoryAuthorRepository{}

	presenter := SendMessageTestPresenter{}

	handler := SendMessageHandler(&authors_repository, &message_repository)

	handler(dto, &presenter)

	assert.Equal(t, "author_not_found", presenter.response)
	assert.Equal(t, 0, len(message_repository.messages))
}

