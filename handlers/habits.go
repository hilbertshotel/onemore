package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"onemore/habit"
	"onemore/logger"
	"strings"
   
 //   "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func habitsHandler(w http.ResponseWriter, r *http.Request, log *logger.Logger, coll *mongo.Collection) {
	w.Header().Set("content-type", "application/json")

	var path []string
	for _, v := range strings.Split(r.URL.Path, "/") {
		if v != "" {
			path = append(path, v)
		}
	}

	// GET HABITS
	if len(path) == 1 && r.Method == http.MethodGet {
		// get all habits from database
		habits, err := habit.Get(coll, log)
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			return
		}

		// marshal data
		out, err := json.Marshal(habits)
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			log.Error(err)
			return
		}

		// return data to frontend
		w.Write(out)
		return
	}

	// PUT HABIT
	if len(path) == 1 && r.Method == http.MethodPut {
		var id primitive.ObjectID

        body, err := ioutil.ReadAll(r.Body)
        if err != nil {
            http.Error(w, "Inernatl Server Error", 500)
            log.Error(err)
            return
        }

        err = json.Unmarshal(body, &id)
        if err != nil {
            http.Error(w, "Internal Server Error", 500)
            log.Error(err)
            return
        }

		err = habit.Increment(id, coll)
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
		// read request body
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			log.Error(err)
			return
		}

		// unmarshal request body
		var newHabitName string
		err = json.Unmarshal(data, &newHabitName)
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			log.Error(err)
			return
		}

		// post data to db and retrieve it
		habit, err := habit.Post(newHabitName, coll, log)
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			return
		}

		// marshal data
		out, err := json.Marshal(habit)
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			log.Error(err)
			return
		}

		// return data to frontend
		w.Write(out)
		msg := fmt.Sprintf("Added new habit: %v", newHabitName)
		log.Ok(msg)
		return
	}

	http.NotFound(w, r)
}
