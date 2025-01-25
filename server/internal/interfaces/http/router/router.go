package router

import (
	"net/http"

	"github.com/mcorrigan89/bigapp/server/internal/interfaces/http/handlers"
	"github.com/mcorrigan89/bigapp/server/internal/interfaces/http/middleware"
)

func NewRouter(mux *http.ServeMux, middleware middleware.Middleware, userHandler *handlers.UserHandler, imageHandler *handlers.ImageHandler) http.Handler {

	// User routes
	mux.HandleFunc("GET /user/{id}", userHandler.GetUserByID)
	mux.HandleFunc("GET /user/email/{email}", middleware.Authorization(userHandler.GetUserByEmail))

	// Image routes
	mux.HandleFunc("POST /image/upload", imageHandler.UploadImage)
	mux.HandleFunc("POST /image/{collectionID}/uploads", imageHandler.UploadImages)
	mux.HandleFunc("GET /image/{id}/metadata", middleware.Authorization(imageHandler.GetImageByID))
	mux.HandleFunc("GET /image/{id}", imageHandler.GetImageDataByID)

	return middleware.RecoverPanic(middleware.EnabledCORS(middleware.ContextBuilder(mux)))
}
