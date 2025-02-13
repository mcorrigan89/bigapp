package service

import (
	"net/http"
	"sync"

	"connectrpc.com/connect"
	"connectrpc.com/grpcreflect"

	"github.com/mcorrigan89/bigapp/server/internal/application"
	"github.com/mcorrigan89/bigapp/server/internal/common"
	mediav1connect "github.com/mcorrigan89/bigapp/server/internal/interfaces/rpc/gen/media/v1/mediav1connect"
	organizationv1connect "github.com/mcorrigan89/bigapp/server/internal/interfaces/rpc/gen/organization/v1/organizationv1connect"
	userv1connect "github.com/mcorrigan89/bigapp/server/internal/interfaces/rpc/gen/user/v1/userv1connect"
	"github.com/mcorrigan89/bigapp/server/internal/interfaces/rpc/handlers"
	"github.com/rs/zerolog"
)

func NewRpcRoutes(mux *http.ServeMux, cfg *common.Config, logger *zerolog.Logger, wg *sync.WaitGroup, userApplicationService application.UserApplicationService, imageApplicationService application.ImageApplicationService, organizationApplicationService application.OrganizationApplicationService) {

	interceptors := connect.WithInterceptors(newAuthInterceptor(cfg.ServerToken))

	userV1Service := handlers.NewUserServiceV1(logger, userApplicationService)
	imageV1Service := handlers.NewImageServiceV1(logger, imageApplicationService)
	organizationV1Service := handlers.NewOrganizationServiceV1(logger, organizationApplicationService)

	reflector := grpcreflect.NewStaticReflector(
		"user.v1.UserService",
		"media.v1.ImageService",
	)

	reflectPath, reflectHandler := grpcreflect.NewHandlerV1(reflector)
	mux.Handle(reflectPath, reflectHandler)
	reflectPathAlpha, reflectHandlerAlpha := grpcreflect.NewHandlerV1Alpha(reflector)
	mux.Handle(reflectPathAlpha, reflectHandlerAlpha)

	userV1Path, userV1Handle := userv1connect.NewUserServiceHandler(userV1Service, interceptors)
	mux.Handle(userV1Path, userV1Handle)
	imageV1Path, imageV1Handle := mediav1connect.NewImageServiceHandler(imageV1Service, interceptors)
	mux.Handle(imageV1Path, imageV1Handle)
	organizationV1Path, organizationV1Handle := organizationv1connect.NewOrganizationServiceHandler(organizationV1Service, interceptors)
	mux.Handle(organizationV1Path, organizationV1Handle)
}
