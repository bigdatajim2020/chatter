package main

import (
	"chatter/datastore"
	"net/http"
)

// errHandler handles GET: /err
func errHandler(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	msg := q.Get("msg")
	_, err := session(w, r)
	if err != nil {
		renderHTML(w, msg, "layout", "public.navbar", "error")
	} else {
		renderHTML(w, msg, "layout", "public.navbar", "error")
	}
}

// indexHandler handles GET: /
func indexHandler(w http.ResponseWriter, r *http.Request) {
	threads, err := datastore.Threads()
	if err != nil {
		errRedirect(w, r, "Cannot get threads")
	} else {
		_, err = session(w, r)
		if err != nil {
			renderHTML(w, threads, "layout", "public.navbar", "index")
		} else {
			renderHTML(w, threads, "layout", "private.navbar", "index")
		}
	}
}
