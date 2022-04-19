package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"onemore/habit"
	"onemore/logger"
	"strconv"
	"strings"
)

func habitsHandler(w http.ResponseWriter, r *http.Request, log *logger.Logger) {
	w.Header().Set("content-type", "application/json")

	var path []string
	for _, v := range strings.Split(r.URL.Path, "/") {
		if v != "" {
			path = append(path, v)
		}
	}

	// GET HABITS
	if len(path) == 1 && r.Method == http.MethodGet {
		habits, err := habit.Get()
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			log.Error(err)
			return
		}

		out, err := json.Marshal(habits)
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			log.Error(err)
			return
		}

		w.Write(out)
		return
	}

	// PUT HABIT
	if len(path) == 2 && r.Method == http.MethodPut {
		id, err := strconv.Atoi(path[1])
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			log.Error(err)
			return
		}

		err = habit.Put(id)
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			log.Error(err)
			return
		}

		msg := fmt.Sprintf("Incremented habit with id: %v", id)
		log.Ok(msg)
		return
	}

	// POST HABIT
	if len(path) == 1 && r.Method == http.MethodPost {
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			log.Error(err)
			return
		}

		var newHabitName string
		err = json.Unmarshal(data, &newHabitName)
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			log.Error(err)
			return
		}

		habit, err := habit.Post(newHabitName)
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			log.Error(err)
			return
		}

		out, err := json.Marshal(habit)
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			log.Error(err)
			return
		}

		w.Write(out)
		msg := fmt.Sprintf("Added new habit: %v", newHabitName)
		log.Ok(msg)
		return
	}

	http.NotFound(w, r)
}
