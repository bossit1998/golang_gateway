package main

import (
	"fmt"

	//"github.com/casbin/casbin/v2"
	//defaultrolemanager "github.com/casbin/casbin/v2/rbac/default-role-manager"
	//"github.com/casbin/casbin/v2/util"
	//gormadapter "github.com/casbin/gorm-adapter/v2"
	"github.com/gomodule/redigo/redis"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"bitbucket.org/alien_soft/api_getaway/api"
	"bitbucket.org/alien_soft/api_getaway/config"
	"bitbucket.org/alien_soft/api_getaway/pkg/grpc_client"
	"bitbucket.org/alien_soft/api_getaway/pkg/logger"
	"bitbucket.org/alien_soft/api_getaway/storage"
	rds "bitbucket.org/alien_soft/api_getaway/storage/redis"
	"bitbucket.org/alien_soft/api_getaway/storage/repo"
)

var (
	log        		logger.Logger
	cfg        		config.Config
	strg       		storage.StorageI
	inMemStrg      repo.InMemoryStorageI
	grpcClient 		*grpc_client.GrpcClient
	//casbinEnforcer 	*casbin.Enforcer
)

func initDeps() {
	cfg = config.Load()
	log = logger.New(cfg.LogLevel, "api_getaway")

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

	/*
	a, err := gormadapter.NewAdapter("postgres", psqlString, true)
	if err != nil {
		log.Error("new adapter error", logger.Error(err))
		return
	}

	casbinEnforcer, err = casbin.NewEnforcer(cfg.CasbinConfigPath, a)
	if err != nil {
		log.Error("new enforcer error", logger.Error(err))
		return
	}

	err = casbinEnforcer.LoadPolicy()
	if err != nil {
		log.Error("casbin load policy error", logger.Error(err))
		return
	}
	*/

	pool := redis.Pool{
		// Maximum number of idle connections in the pool.
		MaxIdle: 80,
		// max number of connections
		MaxActive: 12000,
		// Dial is an application supplied function for creating and
		// configuring a connection.
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", fmt.Sprintf("%s:%d", cfg.RedisHost, cfg.RedisPort))
			if err != nil {
				panic(err.Error())
			}
			return c, err
		},
	}

	inMemStrg = rds.NewRedisRepo(&pool)

	grpcClient, err = grpc_client.New(cfg)
	if err != nil {
		log.Error("grpc dial error", logger.Error(err))
	}
}

func main() {
	initDeps()

	//casbinEnforcer.GetRoleManager().(*defaultrolemanager.RoleManager).AddMatchingFunc("keyMatch", util.KeyMatch)
	//casbinEnforcer.GetRoleManager().(*defaultrolemanager.RoleManager).AddMatchingFunc("KeyMatch3", util.KeyMatch3)

	server := api.New(api.Config{
		Storage:    strg,
		InMemoryStorage: inMemStrg,
		Logger:     log,
		GrpcClient: grpcClient,
	//	CasbinEnforcer:  casbinEnforcer,
		Cfg:        cfg,
	})

	server.Run(cfg.HTTPPort)
}
