package middleware

import (
	"net/http"

	"github.com/mcorrigan89/simple_auth/server/internal/domain/entities"
)

func (m *middleware) Authorization(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		userContext, ok := ctx.Value(currentUserContextKey).(*entities.UserContextEntity)
		if !ok {
			m.logger.Error().Ctx(ctx).Msg("User context not found in request context")
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		if userContext.IsExpired() {
			m.logger.Error().Ctx(ctx).Msg("User context has expired")
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
