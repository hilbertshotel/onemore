package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"onemore/habit"
	"onemore/logger"
	"strings"
    "time"
    "context"

    "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func habitsHandler(w http.ResponseWriter, r *http.Request, log *logger.Logger, coll *mongo.Collection) {
	w.Header().Set("content-type", "application/json")

    // unpack url
	var path []string
	for _, v := range strings.Split(r.URL.Path, "/") {
		if v != "" {
			path = append(path, v)
		}
	}

    // handle wrong url
    if len(path) != 1 {
        http.NotFound(w, r)    
    }

    // create context for db
    ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
    defer cancel()

	// GET
	if r.Method == http.MethodGet {
		// get all habits from database
		habits, err := habit.Get(coll, log, ctx)
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

	// PUT
	if r.Method == http.MethodPut {
		var id primitive.ObjectID
        
        // unpack request data
        body, err := ioutil.ReadAll(r.Body)
        if err != nil {
            http.Error(w, "Inernatl Server Error", 500)
            log.Error(err)
            return
        }

        // unmarshal request data
        err = json.Unmarshal(body, &id)
        if err != nil {
            http.Error(w, "Internal Server Error", 500)
            log.Error(err)
            return
        }

        // increment habit
		err = habit.Increment(id, coll, ctx)
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			log.Error(err)
			return
		}
        
		msg := fmt.Sprintf("Incremented habit with id: %v", id)
		log.Ok(msg)
		return
	}

	// POST
	if r.Method == http.MethodPost {
		// read request body
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			log.Error(err)
			return
		}

		// unmarshal request body
		var name string
		err = json.Unmarshal(data, &name)
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			log.Error(err)
			return
		}

		// post data to db and retrieve it
		habit, err := habit.Post(name, coll, log, ctx)
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
		msg := fmt.Sprintf("Added new habit: %v", name)
		log.Ok(msg)
		return
	}

	http.NotFound(w, r)
}
