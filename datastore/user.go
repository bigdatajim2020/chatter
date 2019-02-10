package datastore

import "time"

// A User represents the forum user's information.
type User struct {
	ID        int
	UUID      string
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
}

// A Session represents a user's current login session.
type Session struct {
	ID        int
	UUID      string // randomly generated unique ID as a cookie value.
	Email     string // User registered email.
	UserID    int    // ID of user table row in database.
	CreatedAt time.Time
}

// UserByEmail gets a single user by the given email.
func UserByEmail(email string) (u User, err error) {
	q := `
		select
			id, uuid, name, email, password, created_at
		from
			users
		where
			email = $1
	`
	err = Db.QueryRow(q, email).Scan(&u.ID, &u.UUID, &u.Name, &u.Email, &u.Password, &u.CreatedAt)
	return
}
