package middleware

import (
	"context"
	"net/http"

	"github.com/rs/xid"
)

type contextKey string

const (
	ipKey            contextKey = "ip"
	correlationIDKey contextKey = "correlation_id"
	sessionTokenKey  contextKey = "sessionTokenKey"
)

// func getCorrelationIdFromContext(ctx context.Context) string {
// 	correlationId, ok := ctx.Value(correlationIDKey).(string)
// 	if !ok {
// 		return ""
// 	}
// 	return correlationId
// }

// func getSessionTokenFromContext(ctx context.Context) string {
// 	sessionToken, ok := ctx.Value(sessionTokenKey).(string)
// 	if !ok {
// 		return ""
// 	}
// 	return sessionToken
// }

// func getIPFromContext(ctx context.Context) string {
// 	ip, ok := ctx.Value(ipKey).(string)
// 	if !ok {
// 		return ""
// 	}
// 	return ip
// }

func (m *middleware) ContextBuilder(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		ctx = context.WithValue(ctx, ipKey, r.RemoteAddr)
		correlationID := xid.New().String()
		ctx = context.WithValue(ctx, correlationIDKey, correlationID)

		// sessionToken := r.Header.Get(identity.SessionTokenKey)
		// ctx = context.WithValue(ctx, sessionTokenKey, sessionToken)

		ctx = m.logger.WithContext(ctx)

		// user, session, err := app.services.identity.UserService.GetUserBySessionToken(ctx, sessionToken)
		// ctx = identity.ContextSetSession(ctx, session)
		// if err == nil && !session.IsExpired() {
		// 	ctx = identity.ContextSetUser(ctx, user)
		// }

		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
