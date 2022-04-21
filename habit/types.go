package habit

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Habit struct {
	Id      primitive.ObjectID `bson:"_id,omitempty"`
	Name    string             `bson:"name"`
	Days    int                `bson:"days"`
	Inc     bool               `bson:"inc"`
	LastInc string             `bson:"last_inc"`
}
