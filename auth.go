package main

import (
	"chatter/datastore"
	"net/http"
)

// authenticate verifies user by email, then password, redirects to login page if credential incorrect.
// TODO: improve authentication logic.
func authenticate(w http.ResponseWriter, r *http.Request) {
	user, err := datastore.UserByEmail(r.PostFormValue("email"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if user.Password == datastore.Encrypt(r.PostFormValue("password")) {
		session, err := user.NewSession()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		cookie := http.Cookie{
			Name:     "_cookie",
			Value:    session.UUID,
			HttpOnly: true,
		}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		http.Redirect(w, r, "/login", http.StatusFound)
	}
}
