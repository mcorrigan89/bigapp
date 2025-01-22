package router

import (
	"net/http"

	"github.com/mcorrigan89/simple_auth/server/internal/interfaces/http/handlers"
	"github.com/mcorrigan89/simple_auth/server/internal/interfaces/http/middleware"
)

func NewRouter(mux *http.ServeMux, middleware middleware.Middleware, userHandler *handlers.UserHandler) http.Handler {

	mux.HandleFunc("GET /user/{id}", userHandler.GetUserByID)
	mux.HandleFunc("GET /user/email/{email}", middleware.Authorization(userHandler.GetUserByEmail))

	return middleware.RecoverPanic(middleware.EnabledCORS(middleware.ContextBuilder(mux)))
}
