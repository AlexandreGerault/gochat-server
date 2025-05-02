package domain

type Message struct {
	content string
}

func NewMessage(content string) Message {
	return Message{content}
}
