package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
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

	errChan := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		errChan <- http.ListenAndServe(":8080", nil)
	}()

	log.Fatal(<-errChan)
}
