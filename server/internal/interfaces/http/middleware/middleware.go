package middleware

import (
	"net/http"

	"github.com/mcorrigan89/simple_auth/server/internal/common"
	"github.com/rs/zerolog"
)

type Middleware interface {
	ContextBuilder(next http.Handler) http.Handler
	RecoverPanic(next http.Handler) http.Handler
	EnabledCORS(next http.Handler) http.Handler
}

type middleware struct {
	config *common.Config
	logger *zerolog.Logger
}

func CreateMiddleware(config *common.Config, logger *zerolog.Logger) *middleware {
	return &middleware{config: config, logger: logger}
}
