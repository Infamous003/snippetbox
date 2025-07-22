package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	// We are defining a command line argument for port
	// using the string function. it converts anything that you enter to a string, if it can that is
	// the string func retuns the address of the addr variable that line sin the runtime
	// default value is 4000
	addr := flag.String("addr", ":4000", "Port number")

	// Then we need to parse it
	flag.Parse()

	// Creating a logger for informational messages and errors
	// the last argument uses bitwise or to sombine date and time of the log
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		infoLog:  infoLog,
		errorLog: errorLog,
	}

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/view", app.snippetView)
	mux.HandleFunc("/snippet/create", app.snippetCreate)

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  mux,
	}

	infoLog.Printf("Starting server on port %s", *addr)

	// Sending the pointer address. deferencing it basiclaly
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
