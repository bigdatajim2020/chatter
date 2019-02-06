package main

import "net/http"

func authenticate(w http.ResponseWriter, r *http.Request) {
	user, err := data.UserByEmail(r.PostFormValue("email"))
	if err != nil {
		//
	}
	if user.Password == data.Encrypt(r.PostFormValue("password")) {
		session := user.CreateSession()
		cookie := http.Cookie{
			Name:     "_cookie",
			Value:    session.Uuid,
			HttpOnly: true,
		}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		http.Redirect(w, r, "/login", http.StatusFound)
	}
}
