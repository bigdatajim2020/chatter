package datastore

// Data for testing.
var users = []User{
	{
		Name:     "Peter Jones",
		Email:    "peter@gmail.com",
		Password: "peter_pass",
	},
	{
		Name:     "John Smith",
		Email:    "john@gmail.com",
		Password: "john_pass",
	},
}

func teardown() (err error) {
	for _, t := range []string{"sessions", "posts", "threads", "users"} {
		err = DeleteAll(t)
		if err != nil {
			return
		}
	}
	return
}
