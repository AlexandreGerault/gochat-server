package domain

import "github.com/google/uuid"

type Message struct {
	Id        uuid.UUID
	Room_Id   uuid.UUID
	Author_Id uuid.UUID
	Content   string
}

func NewMessage(message_id uuid.UUID, room_id uuid.UUID, author_id uuid.UUID, content string) Message {
	return Message{message_id, room_id, author_id, content}
}

type MessageRepository interface {
	Save(message Message) (uuid.UUID, error)
}
