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

// NewSession creates a new session for an existing user when user login is authenticated.
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

// Check checks validation of every request with its cookies against a logged in user. It's used by session util.
func (s *Session) Check() (valid bool, err error) {
	q := `
		select
			id, uuid, email, user_id, created_at
		from
			sessions
		where
			uuid = $1
	`
	err = Db.QueryRow(q, s.UUID).Scan(&s.ID, &s.UUID, &s.Email, &s.UserID, &s.CreatedAt)
	if err != nil {
		return
	}
	if s.ID != 0 {
		valid = true
	}
	return
}

// New creates a new user, save user info into database.
func (u *User) New() (err error) {
	q := `
		insert into
			users (uuid, name, email, password, created_at)
		values ($1, $2, $3, $4, $5)
		returning
			id, uuid, created_at
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
	err = stmt.QueryRow(uuid, u.Name, u.Email, Encrypt(u.Password), time.Now()).Scan(&u.ID, &u.UUID, &u.CreatedAt)
	return
}

// UserByEmail gets a single user by the given email when an existing user attempts to login. It is used in authenticate function.
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
