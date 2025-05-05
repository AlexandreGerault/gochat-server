package http

import (
	"net/http"

	"alexandre-gerault.fr/gochat-server/internal/messaging/application"
	"alexandre-gerault.fr/gochat-server/internal/messaging/domain"
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

func (p *SendMessagePresenter) Presents() {
	p.writer.WriteHeader(http.StatusCreated)
}

func (p *SendMessagePresenter) InvalidPayload() {
	p.writer.WriteHeader(http.StatusBadRequest)
}



func NewSendMessageEndpoint(app *shared_infrastructure.Application) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		handler := application.SendMessageHandler(
			app.Container().Resolve(domain.AuthorRepository),
			app.Container().Resolve(domain.MessageRepository),
			app.Container().Resolve(application.UuidProvider),
		)

		presenter := &SendMessagePresenter{writer}

		handler(
			application.SendMessageDto{
				Room_id: request.FormValue("room_id"),
				Content: request.FormValue("content"),
			},
			presenter,
		)
	}
}
