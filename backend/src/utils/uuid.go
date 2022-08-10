package util

import (
	"github.com/google/uuid"
)

type UUIDGeneratorInterface interface {
	NewRandom() (uuid.UUID, error)
}

// UUIDGenerator is the concrete implementation of the UUIDGeneratorInterface used to
// generate UUIDs in production deployments.
type UUIDGenerator struct {
}

func NewUUIDGenerator() *UUIDGenerator {
	return &UUIDGenerator{}
}

func (r *UUIDGenerator) NewRandom() (uuid.UUID, error) {
	return uuid.NewRandom()
}
