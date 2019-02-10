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

// Threads extracts all threads in the database foe the index handler.
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
