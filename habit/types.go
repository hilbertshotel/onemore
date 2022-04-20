package habit

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Habit struct {
	Id      primitive.ObjectID `bson:"_id,omitempty"`
	Name    string             `bson:"name"`
	Days    int                `bson:"days"`
	Streak  int                `bson:"streak"`
	Inc     bool               `bson:"inc"`
	Active  bool               `bson:"active"`
	LastInc string             `bson:"last_inc"`
	Created string             `bson:"created"`
}
