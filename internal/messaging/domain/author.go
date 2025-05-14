package domain

import "github.com/google/uuid"

type Author struct {
	Uuid uuid.UUID
}

func NewAuthor(uuid uuid.UUID) Author {
	return Author{
		uuid,
	}
}

type AuthorRepository interface {
	Exist(id uuid.UUID) bool
}
