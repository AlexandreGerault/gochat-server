package http

import (
	"fmt"
	"io"
	"net/http"

	"alexandre-gerault.fr/gochat-server/internal/messaging/application"
	shared_infrastructure "alexandre-gerault.fr/gochat-server/internal/shared/infrastructure"
)

type SendMessagePresenter struct {
	writer http.ResponseWriter
}

func (p *SendMessagePresenter) AuthorNotFound() {
	p.writer.WriteHeader(http.StatusNotFound)
}

func (p *SendMessagePresenter) MessageEmpty() {
	p.writer.WriteHeader(http.StatusBadRequest)
}

func (p *SendMessagePresenter) TooLongMessage() {
	p.writer.WriteHeader(http.StatusBadRequest)
}

func (p *SendMessagePresenter) InvalidPayload() {
	p.writer.WriteHeader(http.StatusBadRequest)
}

func (p *SendMessagePresenter) UnexpectedError(error string) {
	p.writer.WriteHeader(http.StatusInternalServerError)
	io.WriteString(p.writer, fmt.Sprintf("{\"message\": \"%s\"}", error))
}

func (p *SendMessagePresenter) MessageSentSuccessfully() {
	p.writer.WriteHeader(http.StatusCreated)
}

func NewSendMessageEndpoint(app *shared_infrastructure.Application) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		handler := application.SendMessageHandler(
			app.Dependencies.Author_Repository,
			app.Dependencies.Message_Repository,
			app.Dependencies.Uuid_Provider,
		)

		presenter := &SendMessagePresenter{writer}

		handler(
			application.SendMessageDto{
				Author_id: request.FormValue("author_id"),
				Room_id:   request.FormValue("room_id"),
				Content:   request.FormValue("content"),
			},
			presenter,
		)
	}
}
