package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Infamous003/snippetbox/internal/models"
	"github.com/jackc/pgx/v5"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	snippets *models.SnippetModel
}

const DATABASE_URL = "postgres://infamous:Getalife@03@localhost:5432/snippetbox"

func main() {
	// Creating a command line argument/flag for port number & parse it
	addr := flag.String("addr", ":4000", "Port number")
	flag.Parse()
	// Creating a logger for informational messages and errors
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	fmt.Println("Connecting to database...")
	conn, err := openDB(DATABASE_URL)
	if err != nil {
		log.Printf("Unable to connect to database: %v\n", err.Error())
	}
	fmt.Println("Database connection established!")
	defer conn.Close(context.Background())

	app := &application{
		infoLog:  infoLog,
		errorLog: errorLog,
		snippets: &models.SnippetModel{DB: conn},
	}

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

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
	conn, err := pgx.Connect(context.Background(), databaseUrl)
	if err != nil {
		return nil, err
	}

	if err = conn.Ping(context.Background()); err != nil {
		return nil, err
	}
	return conn, nil
}
