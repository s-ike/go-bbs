package main

import (
	"net/http"
)

type writer struct{}

func newWriter() *writer {
	return &writer{}
}

func (wr *writer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var msg message
	msg.Name = r.Form["name"][0]
	msg.Body = r.Form["body"][0]
	if msg.Name == "" {
		msg.Name = "名無し"
	}

	err := msg.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusFound)
}
