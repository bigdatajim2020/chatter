package main

import (
	"net/http"
	"chatter/datastore"
)

// newThread handles GET: /new, it shows users an entry form to create a new thread with topic.
func newThread(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
	} else {
		renderHTML(w, nil, "layout", "private.navbar", "new.thread")
	}
}
