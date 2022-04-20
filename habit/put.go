package habit

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Increment(id int, coll *mongo.Collection) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	_, err := coll.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": bson.M{"inc": true}})
	if err != nil {
		return err
	}

	return nil
}

func decrement(habit Habit, coll *mongo.Collection) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	update := bson.M{"$set": bson.M{"inc": false}}
	_, err := coll.UpdateOne(ctx, bson.M{"_id": habit.Id}, update)
	if err != nil {
		return err
	}

	return nil
}

func deactivate(habit Habit, coll *mongo.Collection) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	update := bson.M{"$set": bson.M{"active": false, "days": 0}}
	_, err := coll.UpdateOne(ctx, bson.M{"_id": habit.Id}, update)
	if err != nil {
		return err
	}

	return nil
}
