package shared_infrastructure

import (
	"github.com/google/uuid"
)

type UuidGenerator struct{}

func (uuid_generator *UuidGenerator) Generate() (uuid.UUID, error) {
	uuid, err := uuid.NewV7()

	return uuid, err
}
