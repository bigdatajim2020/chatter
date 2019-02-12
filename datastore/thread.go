package datastore

import (
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
