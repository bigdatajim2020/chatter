package datastore

import (
	"chatter/logger"
	"database/sql"
	"errors"
	"fmt"
	"time"
)

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
	stmt, err := Db.PrepareContext(ctx, q)
	if err != nil {
		return
	}
	defer stmt.Close()

	uuid, err := createUUID()
	if err != nil {
		return
	}
	err = stmt.QueryRow(uuid, u.Email, u.ID, time.Now()).Scan(&s.ID, &s.UUID, &s.Email, &s.UserID, &s.CreatedAt)
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
	err = Db.QueryRowContext(ctx, q, s.UUID).Scan(&s.ID, &s.UUID, &s.Email, &s.UserID, &s.CreatedAt)
	switch {
	case err == sql.ErrNoRows:
		err = errors.New("session with this user not exist")
	case err != nil:
		return
	case s.ID != 0:
		valid = true
	}
	return
}

// GetUser returns a User by session from database.
func (s *Session) GetUser() (u User, err error) {
	q := `
		select
			id, uuid, name, email, created_at
		from
			users
		where
			id = $1
	`
	err = Db.QueryRow(q, s.UserID).Scan(&u.ID, &u.UUID, &u.Name, &u.Email, &u.CreatedAt)
	return
}

// NewThread creates a new thread record with its topic in database.
func (u *User) NewThread(topic string) (t Thread, err error) {
	q := `
		insert into
			threads (uuid, topic, user_id, created_at)
		values ($1, $2, $3, $4)
		returning id, uuid, topic, user_id, created_at
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
	err = stmt.QueryRow(uuid, topic, u.ID, time.Now()).Scan(&t.ID, &t.UUID, &t.Topic, &t.UserID, &t.CreatedAt)
	return
}

// NewPost creates a new post record under a specific thread.
func (u *User) NewPost(t Thread, body string) (p Post, err error) {
	q := `
		insert into
			posts (uuid, body, user_id, thread_id, created_at)
		values ($1, $2, $3, $4, $5)
		returning id, uuid, body, user_id, thread_id, created_at
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
	err = stmt.QueryRow(uuid, body, u.ID, t.ID, time.Now()).Scan(&p.ID, p.UUID, p.Body, p.UserID, p.ThreadID, p.CreatedAt)
	return
}

// DeleteByUUID deletes session record from database when user logs out.
func (s *Session) DeleteByUUID() (err error) {
	q := `
		delete from
			sessions
		where
			uuid = $1
	`
	stmt, err := Db.Prepare(q)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(s.UUID)
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
	stmt, err := Db.PrepareContext(ctx, q)
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
	err = Db.QueryRowContext(ctx, q, email).Scan(&u.ID, &u.UUID, &u.Name, &u.Email, &u.Password, &u.CreatedAt)
	switch {
	case err == sql.ErrNoRows:
		logger.Warning.Printf("No user with email: %s", email)
		err = fmt.Errorf("No user with email: %s", email)
	case err != nil:
		return
	}
	return
}
