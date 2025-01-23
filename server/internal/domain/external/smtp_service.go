package external

import (
	"context"

	"github.com/mcorrigan89/bigapp/server/internal/domain/entities"
)

type SmtpService interface {
	SendEmail(ctx context.Context, email *entities.EmailEntity) error
}
