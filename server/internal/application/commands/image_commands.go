package commands

import (
	"io"
)

type CreateNewImageCommand struct {
	ObjectID string
	File     io.Reader
	Size     int64
}
