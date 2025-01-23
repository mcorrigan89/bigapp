package external

import (
	"github.com/mcorrigan89/bigapp/server/internal/domain/entities"
)

type EmailTemplateService interface {
	LoginEmail(templateFile string, refLink *entities.ReferenceLinkEntity) (*string, *string, error)
}
