package testing

import "github.com/google/uuid"

type FakeUuidProvider struct {
	uuidToReturn uuid.UUID
}

func (provider *FakeUuidProvider) Generate() uuid.UUID {
	return provider.uuidToReturn
}

func (provider *FakeUuidProvider) ChangeNextUuid(nextUuid uuid.UUID) {
	provider.uuidToReturn = nextUuid
}
