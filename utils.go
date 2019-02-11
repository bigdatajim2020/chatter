package main

import (
	"chatter/datastore"
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"net/url"
)

// errRedirect redirects users to a error page.
func errRedirect(w http.ResponseWriter, r *http.Request, msg string) {
	q := url.QueryEscape("msg=" + msg)
	http.Redirect(w, r, "/err?"+q, http.StatusNotFound)
}

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

func renderHTML(w http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, f := range filenames {
		files = append(files, fmt.Sprintf("templ/%s.html", f))
	}

	t := template.Must(template.ParseFiles(files...))
	err := t.ExecuteTemplate(w, "Layout", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
