package main

import (
	"chatter/datastore"
	"errors"
	"net/http"
)

func session(w http.ResponseWriter, r *http.Request) (s datastore.Session, err error) {
	cookie, err := r.Cookie("_cookie")
	if err != nil {
		s = datastore.Session{UUID: cookie.Value}
		if ok, _ := s.Check(); !ok {
			err = errors.New("Invalid session")
		}
	}
	return
}
