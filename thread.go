package main

import (
	"net/http"
)

// newThread handles GET: /thread/new, it shows users an entry form to create a new thread with topic.
func newThread(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
	} else {
		renderHTML(w, nil, "layout", "private.navbar", "new.thread")
	}
}

// createThreadHandler handles POST: /thread/create, it creates a new thread record in the database.
func createThreadHandler(w http.ResponseWriter, r *http.Request) {
	s, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
	} else {
		u, err := s.GetUser()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		topic := r.FormValue("topic")
		if _, err := u.NewThread(topic); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/", http.StatusFound)
	}
}
