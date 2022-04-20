package habit

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Get(coll *mongo.Collection) ([]Habit, error) {
	// create context
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// get all entries in habits collection
	cursor, err := coll.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	var rawData []Habit
	if err = cursor.All(ctx, &rawData); err != nil {
		return nil, err
	}

	// handle expirations
	habits, err := handleExpirations(rawData, coll)
	if err != nil {
		return rawData, err
	}

	// send to frontend
	return habits, nil
}
