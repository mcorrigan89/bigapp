package main

import (
	"net/http"
	"os"
	"sync"

	"github.com/mcorrigan89/simple_auth/server/internal/application"
	"github.com/mcorrigan89/simple_auth/server/internal/common"
	"github.com/mcorrigan89/simple_auth/server/internal/domain/services"
	"github.com/mcorrigan89/simple_auth/server/internal/infrastructure/email"
	"github.com/mcorrigan89/simple_auth/server/internal/infrastructure/postgres"
	"github.com/mcorrigan89/simple_auth/server/internal/infrastructure/postgres/repositories"
	"github.com/mcorrigan89/simple_auth/server/internal/interfaces/http/handlers"
	"github.com/mcorrigan89/simple_auth/server/internal/interfaces/http/middleware"
	"github.com/mcorrigan89/simple_auth/server/internal/interfaces/http/router"
	"github.com/mcorrigan89/simple_auth/server/internal/interfaces/rpc/service"
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
	smtpService := email.NewSmtpService(&cfg)

	userService := services.NewUserService(postgresUserRepository, postgresReferenceLinkRepository)
	emailService := services.NewEmailService(smtpService)

	userApplicationService := application.NewUserApplicationService(db, &wg, &cfg, &logger, userService, emailService)
	userHandler := handlers.NewUserHandler(&logger, userApplicationService)

	mdlwr := middleware.CreateMiddleware(&cfg, db, &logger, userService)
	// HTTP Routes
	httpRoutes := router.NewRouter(mux, mdlwr, userHandler)
	// Connect RPC Routes
	service.NewRpcRoutes(mux, &logger, &wg, userApplicationService)

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
