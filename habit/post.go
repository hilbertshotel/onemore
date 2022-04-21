package habit

import (
	"context"
	"errors"
	"fmt"
	"onemore/logger"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func Post(name string, coll *mongo.Collection, log *logger.Logger) (Habit, error) {
	var habit Habit

    // create ctx
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

    // if name too long
    if len(name) > 10 {
        err := errors.New("name out of bounds")
        log.Error(err)
        return habit, err
    }

	// if habit already exists return error
	count, err := coll.CountDocuments(ctx, bson.M{"name": name})
	if err != nil {
		log.Error(err)
		return habit, err
	}

	if count != 0 {
		msg := fmt.Sprintf("entry already exists (%v)", name)
		err := errors.New(msg)
		log.Error(err)
		return habit, err
	}

	// create new habit
	now := time.Now()

	habit = Habit{
		Name:    name,
		Days:    1,
		Inc:     true,
		Active:  true,
		LastInc: now.Format(time.RFC3339),
		Created: fmt.Sprintf("%v %v %v", now.Year(), now.Month(), now.Day()),
	}

	// insert into collection
	res, err := coll.InsertOne(ctx, habit)
	if err != nil {
		log.Error(err)
		return habit, err
	}

	// return to frontend
	habit.Id = res.InsertedID.(primitive.ObjectID)
	return habit, nil
}
