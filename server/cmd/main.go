package main

import (
	"os"
	"sync"

	"github.com/mcorrigan89/simple_auth/server/internal/common"
	"github.com/rs/zerolog"
)

type services struct {
}

type application struct {
	config   common.Config
	wg       *sync.WaitGroup
	logger   *zerolog.Logger
	services *services
	cache    *common.Cache
}

func main() {
	logger := getLogger()

	logger.Info().Msg("Starting server")

	cfg := common.Config{}
	common.LoadConfig(&cfg)

	db, err := openDBPool(&cfg)
	if err != nil {
		logger.Err(err).Msg("Failed to open database connection")
		os.Exit(1)
	}
	defer db.Close()

	cache := common.CreateCache()
	defer cache.Close()

	wg := sync.WaitGroup{}

	s := services{}

	app := &application{
		wg:       &wg,
		config:   cfg,
		logger:   &logger,
		services: &s,
		cache:    cache,
	}

	err = app.serve()
	if err != nil {
		logger.Err(err).Msg("Failed to start server")
		os.Exit(1)
	}
}
