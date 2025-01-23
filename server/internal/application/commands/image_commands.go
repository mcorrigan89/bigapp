package commands

import (
	"io"

	"github.com/google/uuid"
)

type CreateNewAvatarImageCommand struct {
	UserID   uuid.UUID
	ObjectID string
	File     io.Reader
	Size     int64
}
