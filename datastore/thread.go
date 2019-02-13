package datastore

import (
	"log"
	"time"
)

// A Thread represents a forum thread (a conversation among forum users).
type Thread struct {
	ID        int
	UUID      string
	Topic     string
	UserID    int
	CreatedAt time.Time
}

// A Post represents a post (a message added by a forum user) within a thread.
type Post struct {
	ID        int
	UUID      string
	Body      string
	UserID    int
	ThreadID  int
	CreatedAt time.Time
}

// NumReplies returns the number of posts in a therad.
// It's used in index template as a pipline.
func (t *Thread) NumReplies() (count int, err error) {
	q := `
		select count(*)
		from
			posts
		where
			thread_id = $1
	`
	rows, err := Db.Query(q, t.ID)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&count); err != nil {
			return
		}
	}
	return
}

// User returns the user by a thread.
// It's used in index template as a pipline.
func (t *Thread) User() (u User) {
	q := `
	select
		id, uuid, name, email, created_at
	from
		users
	where
		id = $1
	`
	err := Db.QueryRow(q, t.UserID).Scan(&u.ID, &u.UUID, &u.Name, &u.Email, &u.CreatedAt)
	if err != nil {
		log.Printf("query users by id: %v", err)
	}
	return
}

// User returns the user by a thread.
// It's used in thread templates as piplines.
func (p *Post) User() (u User) {
	q := `
	select
		id, uuid, name, email, created_at
	from
		users
	where
		id = $1
	`
	err := Db.QueryRow(q, p.UserID).Scan(&u.ID, &u.UUID, &u.Name, &u.Email, &u.CreatedAt)
	if err != nil {
		log.Printf("query users by id: %v", err)
	}
	return
}

// CreatedAtDate formats the CreatedAt date to display nicely on the screen
// It's used in index template as a pipline.
func (t *Thread) CreatedAtDate() string {
	return t.CreatedAt.Format("Jan 2, 2006 at 3:04pm")
}

// CreatedAtDate formats the CreatedAt date to display nicely on the screen
// It's used in thread templates as piplines.
func (p *Post) CreatedAtDate() string {
	return p.CreatedAt.Format("Jan 2, 2006 at 3:04pm")
}

// Posts returns all posts belonging to a thread.
// It's used in the thread templates as piplines.
func (t *Thread) Posts() (ps []Post) {
	q := `
		select
			id, uuid, body, user_id, thread_id, created_at
		from
			posts
		where
			thread_id = $1
	`
	rows, err := Db.Query(q, t.ID)
	if err != nil {
		log.Printf("query posts by thread_id: %v", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var p Post
		if err := rows.Scan(&p.ID, &p.UUID, &p.Body, &p.UserID, &p.ThreadID, &p.CreatedAt); err != nil {
			log.Printf("scan posts rows: %v", err)
			return
		}
		ps = append(ps, p)
	}
	return
}

// Threads extracts all threads in the database for the index handler.
func Threads() (threads []Thread, err error) {
	q := `
	select
		id, uuid, topic, user_id, created_at
	from
		threads
	order by
		created_at
	desc`
	rows, err := Db.Query(q)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var t Thread
		if err = rows.Scan(&t.ID, &t.UUID, &t.Topic, &t.UserID, &t.CreatedAt); err != nil {
			return
		}
		threads = append(threads, t)
	}
	if err = rows.Err(); err != nil {
		return
	}
	return
}

// ThreadByUUID gets a thread by its uuid.
func ThreadByUUID(uuid string) (t Thread, err error) {
	q := `
		select
			id, uuid, topic, user_id, created_at
		from
			threads
		where
		uuid = $1
	`
	err = Db.QueryRow(q, uuid).Scan(t.ID, t.UUID, t.Topic, t.UserID, t.CreatedAt)
	return
}
