package habit

import (
	"context"
	"time"

    "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Increment(id primitive.ObjectID, coll *mongo.Collection, ctx context.Context) error {
	t := time.Now().Format(time.RFC3339)
    update := bson.M{"$set": bson.M{"inc": true, "last_inc": t}, "$inc": bson.M{"days": 1}}
	_, err := coll.UpdateOne(ctx, bson.M{"_id": id}, update)
	if err != nil {
		return err
	}

	return nil
}

func decrement(habit Habit, coll *mongo.Collection, ctx context.Context) error {
	update := bson.M{"$set": bson.M{"inc": false}}
	_, err := coll.UpdateOne(ctx, bson.M{"_id": habit.Id}, update)
	if err != nil {
		return err
	}

	return nil
}

func delete(habit Habit, coll *mongo.Collection, ctx context.Context) error {
	_, err := coll.DeleteOne(ctx, bson.M{"_id": habit.Id})
	if err != nil {
		return err
	}

	return nil
}
