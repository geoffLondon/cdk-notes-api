package uuid

import (
	guuid "github.com/google/uuid"
)

type UuidGenerator interface {
	New() string
}

type DefaultUuidGenerator struct {
}

func NewDefaultUuidGenerator() *DefaultUuidGenerator {
	return &DefaultUuidGenerator{}
}

func (uuid *DefaultUuidGenerator) New() string {
	return guuid.New().String()
}
