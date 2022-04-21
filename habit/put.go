package habit

import (
	"context"
	"time"

    "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Increment(id primitive.ObjectID, coll *mongo.Collection) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	t := time.Now().Format(time.RFC3339)
    update := bson.M{"$set": bson.M{"inc": true, "lastInc": t}, "$inc": bson.M{"days": 1}}
	_, err := coll.UpdateOne(ctx, bson.M{"_id": id}, update)
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

func delete(habit Habit, coll *mongo.Collection) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	_, err := coll.DeleteOne(ctx, bson.M{"_id": habit.Id})
	if err != nil {
		return err
	}

	return nil
}
