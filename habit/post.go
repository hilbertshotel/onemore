package habit

import (
	"context"
	"errors"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func Post(name string, coll *mongo.Collection) (Habit, error) {
	// create ctx
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// if habit already exists return error
	var habit Habit
	count, err := coll.CountDocuments(ctx, bson.M{"name": name})
	if err != nil {
		return habit, nil
	}

	if count != 0 {
		msg := fmt.Sprintf("entry already exists (%v)", name)
		return habit, errors.New(msg)
	}

	// create new habit
	now := time.Now()

	habit = Habit{
		Name:    name,
		Days:    1,
		Streak:  1,
		Inc:     true,
		Active:  true,
		LastInc: now.Format(time.RFC3339),
		Created: fmt.Sprintf("%v %v %v", now.Year(), now.Month(), now.Day()),
	}

	// insert into collection
	res, err := coll.InsertOne(ctx, habit)
	if err != nil {
		return habit, err
	}

	// return to frontend
	habit.Id = res.InsertedID.(primitive.ObjectID)
	return habit, nil
}
