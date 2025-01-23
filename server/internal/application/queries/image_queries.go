package queries

import (
	"github.com/google/uuid"
)

type ImageByIDQuery struct {
	ID uuid.UUID
}

type ImageDataByIDQuery struct {
	ID uuid.UUID
}
