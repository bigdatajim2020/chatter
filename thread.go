package main

import (
	"chatter/datastore"
	"chatter/logger"
	"fmt"
	"net/http"
)

// newThreadHandler handles GET: /thread/new, it shows users an entry form to create a new thread with topic.
func newThreadHandler(w http.ResponseWriter, r *http.Request) {
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
			logger.Error.Printf("get user by session: %v", err)
			http.Error(w, "Get user error, please try again", http.StatusInternalServerError)
			return
		}

		topic := r.FormValue("topic")
		if _, err := u.NewThread(topic); err != nil {
			logger.Error.Printf("create new thread: %v", err)
			http.Error(w, "Create new thread error, please try again", http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

// readThreadHandler handles GET: /thread/read, it loads a specific thread by uuid.
func readThreadHandler(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	uuid := q.Get("id")
	thread, err := datastore.ThreadByUUID(uuid)
	if err != nil {
		logger.Error.Printf("load thread by uuid: %v", err)
		errRedirect(w, r, "Can't load thread")
	} else {
		_, err := session(w, r)
		if err != nil {
			renderHTML(w, &thread, "layout", "public.navbar", "public.thread")
		} else {
			renderHTML(w, &thread, "layout", "private.navbar", "private.thread")
		}
	}
}

// postThreadHandler handles POST: /thread/post, it creates a post under a specific thread.
func postThreadHandler(w http.ResponseWriter, r *http.Request) {
	s, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
	} else {
		u, err := s.GetUser()
		if err != nil {
			logger.Error.Printf("get user by session: %v", err)
			http.Error(w, "Get user error, please try again", http.StatusInternalServerError)
			return
		}
		body, uuid := r.FormValue("body"), r.FormValue("uuid")
		thread, err := datastore.ThreadByUUID(uuid)
		if err != nil {
			logger.Error.Printf("load thread by uuid: %v", err)
			errRedirect(w, r, "Can't load thread")
			return
		}
		if _, err := u.NewPost(thread, body); err != nil {
			logger.Error.Printf("create thread post: %v", err)
			errRedirect(w, r, "Can't create post")
			return
		}
		http.Redirect(w, r, fmt.Sprintf("/thread/read?id=%s", uuid), http.StatusFound)
	}
}
