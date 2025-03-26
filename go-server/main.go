package main

import (
	"github.com/vitingr/url-shortner/config"
	db "github.com/vitingr/url-shortner/internal/database"
	route "github.com/vitingr/url-shortner/internal/routes"
	"go.uber.org/zap"
)

func main() {
	logger := zap.Must(zap.NewProduction())
	defer logger.Sync()

	secrets, err := config.LoadConfig()
	if err != nil {
		logger.Error(
			"error loading secrets",
			zap.Error(err),
		)
	}

	redisClient, err := db.NewRedisClient(secrets)
	if err != nil {
		logger.Panic("redis failed to initialize",
			zap.Error(err),
		)
	} else {
		logger.Info("redis is connected")
	}

	pgClient, err := db.NewPostgresClient(secrets)
	if err != nil {
		logger.Panic("postgres failed to initialize",
			zap.Error(err),
		)
	} else {
		logger.Info("postgres is connected")
	}

	r := route.SetupRouter(redisClient, pgClient)
	r.Run(":8080")
}
