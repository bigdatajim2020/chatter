package main

import (
	"chatter/datastore"
	"html/template"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	threads, err := datastore.Threads()
	if err != nil {
		//
	}
	publicTempls := []string{
		"templ/layout.html",
		"templ/pubnavbar.html",
		"templ/index.html",
	}
	privateTempls := []string{
		"templ/layout.html",
		"templ/privnavbar.html",
		"templ/index.html",
	}

	var t *template.Template
	_, err = session(w, r)
	if err != nil {
		t = template.Must(template.ParseFiles(privateTempls...))
	} else {
		t = template.Must(template.ParseFiles(publicTempls...))
	}
	t.ExecuteTemplate(w, "layout", threads)
}
