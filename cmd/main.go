package main

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"bitbucket.org/alien_soft/api_gateway/api"
	"bitbucket.org/alien_soft/api_gateway/config"
	"bitbucket.org/alien_soft/api_gateway/pkg/grpc_client"
	"bitbucket.org/alien_soft/api_gateway/pkg/logger"
	"bitbucket.org/alien_soft/api_gateway/storage"
)

var (
	log        logger.Logger
	cfg        config.Config
	strg       storage.StorageI
	grpcClient *grpc_client.GrpcClient
)

func initDeps() {
	cfg = config.Load()
	log = logger.New(cfg.LogLevel, "api_gateway")

	log.Info("main: sqlxConfig",
		logger.String("host", cfg.PostgresHost),
		logger.Int("port", cfg.PostgresPort),
		logger.String("database", cfg.PostgresDatabase))

	psqlString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresDatabase,
	)

	connDB, err := sqlx.Connect("postgres", psqlString)
	strg = storage.NewStoragePg(connDB)

	grpcClient, err = grpc_client.New(cfg)
	if err != nil {
		log.Error("grpc dial error", logger.Error(err))
	}
}

func main() {
	initDeps()

	server := api.New(api.Config{
		Storage:    strg,
		Logger:     log,
		GrpcClient: grpcClient,
		Cfg:        cfg,
	})

	server.Run(cfg.HTTPPort)
}
