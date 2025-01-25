package main

import (
	"net/http"
	"os"
	"sync"

	"github.com/mcorrigan89/bigapp/server/internal/application"
	"github.com/mcorrigan89/bigapp/server/internal/common"
	"github.com/mcorrigan89/bigapp/server/internal/domain/services"
	"github.com/mcorrigan89/bigapp/server/internal/infrastructure/email"
	"github.com/mcorrigan89/bigapp/server/internal/infrastructure/media"
	"github.com/mcorrigan89/bigapp/server/internal/infrastructure/postgres"
	"github.com/mcorrigan89/bigapp/server/internal/infrastructure/postgres/repositories"
	"github.com/mcorrigan89/bigapp/server/internal/infrastructure/storage"
	"github.com/mcorrigan89/bigapp/server/internal/interfaces/http/handlers"
	"github.com/mcorrigan89/bigapp/server/internal/interfaces/http/middleware"
	"github.com/mcorrigan89/bigapp/server/internal/interfaces/http/router"
	"github.com/mcorrigan89/bigapp/server/internal/interfaces/rpc/service"
	"github.com/rs/zerolog"
)

type appServer struct {
	config *common.Config
	wg     *sync.WaitGroup
	logger *zerolog.Logger
}

func main() {
	logger := getLogger()

	logger.Info().Msg("Starting server")

	cfg := common.Config{}
	common.LoadConfig(&cfg)

	db, err := postgres.OpenDBPool(&cfg)
	if err != nil {
		logger.Err(err).Msg("Failed to open database connection")
		os.Exit(1)
	}
	defer db.Close()

	wg := sync.WaitGroup{}
	mux := http.NewServeMux()

	postgresUserRepository := repositories.NewPostgresUserRepository()
	postgresReferenceLinkRepository := repositories.NewPostgresReferenceLinkRepository()
	postgresImageRepository := repositories.NewPostgresImageRepository()
	blobStorageService := storage.NewBlobStorageService(&cfg)
	smtpService := email.NewSmtpService(&cfg)
	imageMediaService := media.NewImageMediaService(blobStorageService)

	userService := services.NewUserService(postgresUserRepository, postgresReferenceLinkRepository, postgresImageRepository)
	emailService := services.NewEmailService(smtpService)
	emailTemplateService := services.NewEmailTemplateService(&cfg)
	imageService := services.NewImageService(postgresImageRepository)

	userApplicationService := application.NewUserApplicationService(db, &wg, &cfg, &logger, userService, emailService, emailTemplateService)
	imageApplicationService := application.NewImageApplicationService(db, &wg, &cfg, &logger, imageService, userService, imageMediaService)
	userHandler := handlers.NewUserHandler(&logger, userApplicationService)
	imageHandler := handlers.NewImageHandler(&logger, imageApplicationService)

	mdlwr := middleware.CreateMiddleware(&cfg, db, &logger, userService)

	// Connect RPC Routes
	service.NewRpcRoutes(mux, &cfg, &logger, &wg, userApplicationService, imageApplicationService)
	// HTTP Routes
	httpRoutes := router.NewRouter(mux, mdlwr, userHandler, imageHandler)

	server := &appServer{
		wg:     &wg,
		config: &cfg,
		logger: &logger,
	}

	err = server.serve(httpRoutes)
	if err != nil {
		logger.Err(err).Msg("Failed to start server")
		os.Exit(1)
	}
}
