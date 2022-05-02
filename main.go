package main

import (
    "fmt"
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

	// init config
	cfg := config.Init()

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

	// init server
	server := http.Server{
		Addr:    cfg.HostAddr,
		Handler: handlers.Mux(log, cfg, coll),
	}

	// start server
	ch := make(chan error)

	go func(ch chan error) {
        fmt.Println()
        log.Ok("API START: " + cfg.HostAddr)
		ch <- server.ListenAndServe()
	}(ch)

	log.Error(<-ch)
}
