package common

import (
	"github.com/google/uuid"
)

func NewID() string {
	// UUID v4
	id, _ := uuid.NewRandom()
	return id.String()
}
