package habit

import (
	"time"
    "context"
    
	"go.mongodb.org/mongo-driver/mongo"
)

func handleExpirations(rawData []Habit, coll *mongo.Collection, ctx context.Context) ([]Habit, error) {
	habits := []Habit{}
	now := time.Now()

	for _, habit := range rawData {
        lastInc, err := time.Parse(time.RFC3339, habit.LastInc)
        if err != nil {
            return nil, err
        }
        
        exp := lastInc.Add(time.Hour * 24)
        if now.Equal(exp) {
            habit.Inc = false
            err := decrement(habit, coll, ctx)
            if err != nil {
                return nil, err
            }
        }
		
        if now.After(exp) {
            err := delete(habit, coll, ctx)
			if err != nil {
				return nil, err
			}
			continue
		}

		habits = append(habits, habit)
	}

	return habits, nil
}
