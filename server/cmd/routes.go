package main

import (
	"net/http"
)

func (app *application) ping(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	app.logger.Info().Ctx(ctx).Msg("/ping")
	w.Write([]byte("OK"))
}

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /ping", app.ping)
	mux.HandleFunc("GET /health", app.ping)
	mux.HandleFunc("GET /ready", app.ping)

	return app.recoverPanic(app.enabledCORS(app.contextBuilder(mux)))
}
