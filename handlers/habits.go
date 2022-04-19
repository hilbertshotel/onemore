package handlers

import (
    "net/http"
    "onemore/logger"
    "strings"
    "encoding/json"
)

type Habit struct {
    Id int
    Name string
    Days int
    Inc bool
}

func habitsHandler(w http.ResponseWriter, r *http.Request, log *logger.Logger) {
    path := strings.Split(r.URL.Path, "/")[1:]

    // GET HABITS
    if len(path) == 1 && r.Method == http.MethodGet {
        habits := []Habit{
            {1, "nodrugs", 10, true},
            {2, "sleep", 3, false},
            {3, "code", 7, true},
        }

        out, err := json.Marshal(habits)
        if err != nil {
            log.Error(err)
            return
        }

        w.Header().Set("content-type", "application/json")
        w.Write(out)
    }
}
