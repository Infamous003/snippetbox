package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

const DATABASE_URL = "postgres://infamous:Getalife@03@localhost:5432/snippetbox"

func main() {
	// Creating a logger for informational messages and errors
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		infoLog:  infoLog,
		errorLog: errorLog,
	}
	// Creating a command line argument/flag for port number
	addr := flag.String("addr", ":4000", "Port number")
	// Then we need to parse it
	flag.Parse()

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	app.infoLog.Println("Connecting to database...")
	conn, err := openDB(DATABASE_URL)
	if err != nil {
		app.errorLog.Printf("Unable to connect to database: %v\n", err.Error())
	}
	app.infoLog.Println("Database connection established!")

	defer conn.Close(context.Background())

	output, err := conn.Exec(context.Background(), "SELECT * FROM snippets")
	if err != nil {
		app.errorLog.Println(err.Error())
	}
	fmt.Println(output)

	infoLog.Printf("Starting server on port %s", *addr)

	// Sending the pointer address. deferencing it basiclaly
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}

func openDB(databaseUrl string) (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), DATABASE_URL)
	if err != nil {
		return nil, err
	}

	if err = conn.Ping(context.Background()); err != nil {
		return nil, err
	}
	return conn, nil
}
