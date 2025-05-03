package domain

import "github.com/google/uuid"

type Message struct {
	id uuid.UUID
	room_id uuid.UUID
	author_id uuid.UUID
	content string
}

func NewMessage(message_id uuid.UUID, room_id uuid.UUID, author_id uuid.UUID, content string) Message {
	return Message{message_id, room_id, author_id, content}
}
