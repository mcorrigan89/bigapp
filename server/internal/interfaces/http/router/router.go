package router

import (
	"net/http"

	"github.com/mcorrigan89/simple_auth/server/internal/interfaces/http/handlers"
	"github.com/mcorrigan89/simple_auth/server/internal/interfaces/http/middleware"
)

func NewRouter(middleware middleware.Middleware, userHandler *handlers.UserHandler) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /user/{id}", userHandler.GetUserByID)
	mux.HandleFunc("GET /user/email/{email}", userHandler.GetUserByEmail)

	return middleware.RecoverPanic(middleware.EnabledCORS(middleware.ContextBuilder(mux)))
}
