package handlers

import (
	"net/http"
	"onemore/config"
	"onemore/logger"
)

func Mux(log *logger.Logger, cfg *config.Config) *http.ServeMux {
	mux := http.NewServeMux()

	frontend := http.FileServer(http.Dir(cfg.Frontend))
	mux.Handle("/", frontend)

	mux.HandleFunc("/habits/", func(w http.ResponseWriter, r *http.Request) {
		habitsHandler(w, r, log)
	})

	return mux
}
