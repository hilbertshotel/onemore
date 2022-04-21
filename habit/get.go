package habit

import (
	"context"
	"onemore/logger"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Get(coll *mongo.Collection, log *logger.Logger) ([]Habit, error) {
	// create context
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// get all entries in habits collection
	cursor, err := coll.Find(ctx, bson.M{})
	if err != nil {
		log.Error(err)
		return nil, err
	}

	var rawData []Habit
	if err = cursor.All(ctx, &rawData); err != nil {
		log.Error(err)
		return nil, err
	}

	// handle expirations
	habits, err := handleExpirations(rawData, coll)
	if err != nil {
		log.Error(err)
		return rawData, err
	}

	// send to frontend
	return habits, nil
}
