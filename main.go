package main

import (
	"code.stakefish.test/service/ip_validator/config"
	"code.stakefish.test/service/ip_validator/pkg/client"
	"code.stakefish.test/service/ip_validator/pkg/handler"
	"code.stakefish.test/service/ip_validator/pkg/repository"
	"code.stakefish.test/service/ip_validator/pkg/server"
	"code.stakefish.test/service/ip_validator/pkg/service"
	"context"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx := context.Background()
	defer ctx.Done()

	cfg, err := config.New()
	if err != nil {
		panic(err)
	}

	// setup DB connection
	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.MongoDBURL))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = mongoClient.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	err = mongoClient.Ping(ctx, &readpref.ReadPref{})
	if err != nil {
		panic(err)
	}

	db := mongoClient.Database(cfg.DBName)

	repoQueries := repository.NewQueries(db)
	resolverClient := client.NewDNSResolver()
	lookupService := service.NewLookupService(resolverClient, repoQueries)

	handlers := handler.NewHandler(lookupService)

	// start web server
	srv := new(server.Server)
	go func() {
		if err := srv.Run(cfg.Listen, handlers.InitRoutes()); err != nil {
			log.Fatal().Msgf("error occurred while running http server: %s", err.Error())
		}
	}()
	log.Info().Msg("web server started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	log.Info().Msgf("web server shutting down")
	if err := srv.Shutdown(context.Background()); err != nil {
		log.Error().Msgf("error occurred on web server shutting down: %s", err.Error())
	}
}
