package habit

import (
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

func handleExpirations(rawData []Habit, coll *mongo.Collection) ([]Habit, error) {
	habits := []Habit{}
	now := time.Now()

	for _, habit := range rawData {
		if !habit.Inc {
			lastInc, err := time.Parse(time.RFC3339, habit.LastInc)
			if err != nil {
				return habits, err
			}

			incExpires := lastInc.Add(time.Hour * 24).Day()
			habitExpires := lastInc.Add(time.Hour * 24 * 2).Day()

			if now.Day() == incExpires {
				habit.Inc = false
				err := decrement(habit, coll)
				if err != nil {
					return nil, err
				}

			} else if now.Day() >= habitExpires {
				err := delete(habit, coll)
				if err != nil {
					return nil, err
				}
				continue
			}
		}

		habits = append(habits, habit)
	}

	return habits, nil
}
