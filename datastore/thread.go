package datastore

import "time"

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
