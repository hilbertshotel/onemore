package handlers

import (
    "onemore/logger"
    "onemore/config"
    "net/http"
)

func Mux(log *logger.Logger, cfg *config.Config) *http.ServeMux {
    mux := http.NewServeMux()

    frontend := http.FileServer(http.Dir(cfg.Frontend))
    mux.Handle("/", frontend)

    return mux
}
