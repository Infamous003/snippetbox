package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	// this is the actual server that serves our files
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	// handler. when you go to /static/ the above line handles it
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/view", app.snippetView)
	mux.HandleFunc("/snippet/create", app.snippetCreate)

	return mux
}
