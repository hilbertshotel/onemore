package handlers

import (
	"net/http"
	"onemore/config"
	"onemore/logger"

	"go.mongodb.org/mongo-driver/mongo"
)

func Mux(log *logger.Logger, cfg *config.Config, coll *mongo.Collection) *http.ServeMux {
	mux := http.NewServeMux()

	frontend := http.FileServer(http.Dir(cfg.Frontend))
	mux.Handle("/", frontend)

	mux.HandleFunc("/habits/", func(w http.ResponseWriter, r *http.Request) {
		habitsHandler(w, r, log, coll)
	})

	return mux
}
