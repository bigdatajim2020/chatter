package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	mux := &http.ServeMux{} // http.NewServeMux()

	fs := http.FileServer(http.Dir("assets/"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/err", errHandler)
	mux.HandleFunc("/login", loginHandler)
	mux.HandleFunc("/logout", logoutHandler)
	mux.HandleFunc("/signup", signupHandler)
	mux.HandleFunc("/signup_account", signupAccountHandler)
	mux.HandleFunc("/authenticate", authenticateHandler)
	mux.HandleFunc("/thread/new", newThreadHandler)
	mux.HandleFunc("/thread/create", createThreadHandler)
	mux.HandleFunc("/thread/post", postThreadHandler)
	mux.HandleFunc("/thread/read", readThreadHandler)

	srv := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      mux,
	}

	errChan := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		errChan <- srv.ListenAndServe()
	}()

	log.Fatal(<-errChan)
}
