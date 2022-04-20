package main

import (
	"context"
	"net/http"
	"onemore/config"
	"onemore/handlers"
	"onemore/logger"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	// init logger
	log := logger.Init()
	log.Ok("Logger initiated")

	// init config
	cfg := config.Init()
	log.Ok("Config initiated")

	// connect to database
	ctx, cancel := context.WithTimeout(context.Background(), cfg.Mongo.Timeout)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.Mongo.Uri))
	if err != nil {
		log.Error(err)
		return
	}
	defer client.Disconnect(ctx)

	coll := client.Database(cfg.Mongo.Database).Collection(cfg.Mongo.Coll)
	log.Ok("Database connection established")

	// init server
	server := http.Server{
		Addr:    cfg.HostAddr,
		Handler: handlers.Mux(log, cfg, coll),
	}

	// start server
	ch := make(chan error)

	go func(ch chan error) {
		log.Ok("Server listening on " + cfg.HostAddr)
		ch <- server.ListenAndServe()
	}(ch)

	log.Error(<-ch)
}
