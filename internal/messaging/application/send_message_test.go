package application

import (
	"fmt"
	"testing"

	"alexandre-gerault.fr/gochat-server/internal/messaging/domain"
	testUtils "alexandre-gerault.fr/gochat-server/internal/testing"
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
	authors []domain.Author
}

func (inMemoryUserRepository *InMemoryAuthorRepository) GetById(id uuid.UUID) (domain.Author, error) {
	for _, author := range inMemoryUserRepository.authors {
		if author.Uuid.String() == id.String() {
			return author, nil
		}
	}

	return domain.Author{}, fmt.Errorf("Cannot find author %s", id.String())
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

func (presenter *SendMessageTestPresenter) InvalidPayload() {
	presenter.response = "invalid_payload"
}

func (presenter *SendMessageTestPresenter) UnexpectedError(error string) {
	presenter.response = error
}

func TestItCanSendMessage(t *testing.T) {
	fakeUuidProvider := testUtils.FakeUuidProvider{}

	author_id, author_err := uuid.Parse("01968e00-1b4d-7a91-bb2a-c55bd56a2dac")
	room_id, room_err := uuid.Parse("01969529-0a14-7556-a125-e9224be7b3ab")

	if author_err != nil || room_err != nil {
		t.Error("Error while parsing a uuid")
	}

	dto := SendMessageDto{author_id.String(), room_id.String(), "Some message"}

	message_repository := InMemoryMessageRepository{}
	authors_repository := InMemoryAuthorRepository{
		[]domain.Author{domain.NewAuthor(author_id)},
	}

	presenter := SendMessageTestPresenter{}

	handler := SendMessageHandler(&authors_repository, &message_repository, &fakeUuidProvider)

	handler(dto, &presenter)

	assert.Equal(t, "success", presenter.response)
	assert.Equal(t, 1, len(message_repository.messages))
}

func TestItCannotSendAnEmptyMessage(t *testing.T) {
	fakeUuidProvider := testUtils.FakeUuidProvider{}

	author_id, author_err := uuid.Parse("01968e00-1b4d-7a91-bb2a-c55bd56a2dac")
	room_id, room_err := uuid.Parse("01969529-0a14-7556-a125-e9224be7b3ab")

	if author_err != nil || room_err != nil {
		t.Error("Error while parsing a uuid")
	}

	dto := SendMessageDto{author_id.String(), room_id.String(), ""}

	message_repository := InMemoryMessageRepository{}
	authors_repository := InMemoryAuthorRepository{
		[]domain.Author{domain.NewAuthor(author_id)},
	}

	presenter := SendMessageTestPresenter{}

	handler := SendMessageHandler(&authors_repository, &message_repository, &fakeUuidProvider)

	handler(dto, &presenter)

	assert.Equal(t, "empty", presenter.response)
	assert.Equal(t, 0, len(message_repository.messages))
}

func TestItCannotSendAnOversizedMessage(t *testing.T) {
	fakeUuidProvider := testUtils.FakeUuidProvider{}

	author_id, auth_err := uuid.Parse("01968e00-1b4d-7a91-bb2a-c55bd56a2dac")
	room_id, room_err := uuid.Parse("01969529-0a14-7556-a125-e9224be7b3ab")

	if auth_err != nil || room_err != nil {
		t.Error("Cannot generate author_id")
	}

	dto := SendMessageDto{author_id.String(), room_id.String(), random.RandString(2001)}

	message_repository := InMemoryMessageRepository{}
	authors_repository := InMemoryAuthorRepository{
		[]domain.Author{domain.NewAuthor(author_id)},
	}

	presenter := SendMessageTestPresenter{}

	handler := SendMessageHandler(&authors_repository, &message_repository, &fakeUuidProvider)

	handler(dto, &presenter)

	assert.Equal(t, "too_long", presenter.response)
	assert.Equal(t, 0, len(message_repository.messages))
}

func TestItCannotSendMessageIfAuthorDoesNotExist(t *testing.T) {
	fakeUuidProvider := testUtils.FakeUuidProvider{}

	author_id, auth_err := uuid.Parse("01968e00-1b4d-7a91-bb2a-c55bd56a2dac")
	room_id, room_err := uuid.Parse("01969529-0a14-7556-a125-e9224be7b3ab")

	if auth_err != nil || room_err != nil {
		t.Error("Cannot generate author_id")
	}

	dto := SendMessageDto{author_id.String(), room_id.String(), "Some legal content"}

	message_repository := InMemoryMessageRepository{}
	authors_repository := InMemoryAuthorRepository{}

	presenter := SendMessageTestPresenter{}

	handler := SendMessageHandler(&authors_repository, &message_repository, &fakeUuidProvider)

	handler(dto, &presenter)

	assert.Equal(t, "author_not_found", presenter.response)
	assert.Equal(t, 0, len(message_repository.messages))
}
