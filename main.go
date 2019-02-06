package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("assets/"))

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/err", errHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/logout", logoutHandler)
	http.HandleFunc("/signup", signupHandler)
	http.HandleFunc("/signup_account", signupAccountHandler)
	http.HandleFunc("/authenticate", authenticateHandler)
	http.HandleFunc("/thread/new", newThreadHandler)
	http.HandleFunc("/thread/create", createThreadHandler)
	http.HandleFunc("/thread/post", postThreadHandler)
	http.HandleFunc("/thread/read", readThreadHandler)

	http.Handle("/static/", http.StripPrefix("/static/", fs))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
