package main

import (
	"chatter/datastore"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	threads, err := datastore.Threads()
	if err != nil {
		//
	}

	_, err = session(w, r)
	if err != nil {
		html(w, threads, "layout", "public.navbar", "index")
	} else {
		html(w, threads, "layout", "private.navbar", "index")
	}
}
