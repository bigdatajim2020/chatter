package main

import (
	"log"
	"net/http"

	"golang.org/x/crypto/acme/autocert"
)

func main() {
	fs := http.FileServer(http.Dir("assets/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

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

	// log.Fatal(http.ListenAndServe(":8080", nil))
	// create the autocert.Manager with domains and path to the cache.
	m := &autocert.Manager{
		Cache:      autocert.DirCache("secret-dir"),
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist("hostname"), // must not be localhost
	}
	s := http.Server{
		Addr:      ":https",
		TLSConfig: m.TLSConfig(),
	}
	go func() {
		// server HTTP, which will redirect automatically to HTTPS.
		h := m.HTTPHandler(nil)
		log.Fatal(http.ListenAndServe(":http", h))
	}()
	// serve HTTPS
	log.Fatal(s.ListenAndServeTLS("", ""))
}
