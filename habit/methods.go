package habit

func Post(name string) (Habit, error) {
	// create new habit struct
	// if habit already exists return error

	// add to database
	// return to frontend
	return Habit{}, nil
}

func Put(id int) error {
	// change incremented status to true for habit with id
	return nil
}

func Get() ([]Habit, error) {
	// get habits out of database
	habits := []Habit{
		{1, "nodrugs", 10, true},
		{2, "sleep", 3, false},
		{3, "code", 7, true},
	}

	return habits, nil
}
