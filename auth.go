package main

import (
	"chatter/datastore"
	"chatter/logger"
	"net/http"
)

// loginHandler handles GET: /login, returns a login page.
func loginHandler(w http.ResponseWriter, r *http.Request) {
	renderHTML(w, nil, "login.layout", "login")
}

// signupHandler handles GET: /signup, returns a signup page.
func signupHandler(w http.ResponseWriter, r *http.Request) {
	renderHTML(w, nil, "login.layout", "signup")
}

// signupAccountHandler handles POST: /signup, this creates an account.
//
// TODO: improve signup logic.
func signupAccountHandler(w http.ResponseWriter, r *http.Request) {
	user := datastore.User{
		Name:     r.FormValue("name"),
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}
	logger.Info.Printf("user: %#v\n", user)
	if err := user.New(); err != nil {
		logger.Error.Printf("create user: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	http.Redirect(w, r, "/login", http.StatusFound)
}

// authenticate handles user login form, POST: /authenticate.
// It verifies user by email, then password, redirecting to login page if credential is incorrect.
//
// TODO: improve authentication logic.
func authenticateHandler(w http.ResponseWriter, r *http.Request) {
	user, err := datastore.UserByEmail(r.PostFormValue("email"))
	if err != nil {
		logger.Warning.Printf("error retrive user by email: %v", err)
		errRedirect(w, r, err.Error())
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

// logoutHandler handles GET: /logout, it logs the user out.
func logoutHandler(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("_cookie")
	if err != http.ErrNoCookie {
		s := datastore.Session{UUID: c.Value}
		s.DeleteByUUID()
		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		// TODO: improve no cookie case.
		http.Error(w, err.Error(), http.StatusExpectationFailed)
	}
}
