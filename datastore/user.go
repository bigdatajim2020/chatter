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

// NewSession creates a new session for an existing user.
func (u *User) NewSession() (s Session, err error) {
	q := `
		insert into
			sessions (uuid, email, user_id, created_at)
		values
			($1, $2, $3, $4)
		returning
			id, uuid, email, user_id, created_at
	`
	stmt, err := Db.Prepare(q)
	if err != nil {
		return
	}
	defer stmt.Close()

	uuid, err := createUUID()
	if err != nil {
		return
	}
	err = stmt.QueryRow(uuid, u.Email, u.ID, time.Now()).Scan(&s.CreatedAt, s.UUID, s.Email, s.UserID, s.CreatedAt)
	return
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
