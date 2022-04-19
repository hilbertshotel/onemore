package main

import (
    "net/http"
    "onemore/config"
    "onemore/logger"
    "onemore/handlers"
)

func main() {
    
    // init logger
    log := logger.Init()
    log.Ok("Logger initiated.")

    // init config
    cfg := config.Init()
    log.Ok("Config initiated.")

    // init server
    server := http.Server{
        Addr: cfg.HostAddr,
        Handler: handlers.Mux(log, cfg),
    }

    // start server
    ch := make(chan error)

    go func(ch chan error) {
        log.Ok("Server listening on " + cfg.HostAddr)
        ch<-server.ListenAndServe()
    }(ch)

    log.Error(<-ch)
}
