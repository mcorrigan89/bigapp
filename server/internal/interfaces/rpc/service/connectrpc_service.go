package service

import (
	"net/http"
	"sync"

	"connectrpc.com/grpcreflect"

	"github.com/mcorrigan89/bigapp/server/internal/application"
	userv1connect "github.com/mcorrigan89/bigapp/server/internal/interfaces/rpc/gen/user/v1/userv1connect"
	"github.com/mcorrigan89/bigapp/server/internal/interfaces/rpc/handlers"
	"github.com/rs/zerolog"
)

func NewRpcRoutes(mux *http.ServeMux, logger *zerolog.Logger, wg *sync.WaitGroup, userApplicationService application.UserApplicationService) {

	userV1Service := handlers.NewUserServiceV1(logger, userApplicationService)

	reflector := grpcreflect.NewStaticReflector(
		"user.v1.UserService",
	)

	reflectPath, reflectHandler := grpcreflect.NewHandlerV1(reflector)
	mux.Handle(reflectPath, reflectHandler)
	reflectPathAlpha, reflectHandlerAlpha := grpcreflect.NewHandlerV1Alpha(reflector)
	mux.Handle(reflectPathAlpha, reflectHandlerAlpha)

	userV1Path, userV1Handle := userv1connect.NewUserServiceHandler(userV1Service)
	mux.Handle(userV1Path, userV1Handle)
}
