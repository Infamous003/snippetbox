package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Home page"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	strId := r.URL.Query().Get("id")
	// Converting the id from str to int
	id, err := strconv.Atoi(strId)

	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Displaying a specific snippet with id:%d", id)
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		w.Header().Set("Allow", "POST")
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("Creating a snipper"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Println("Starting server on port 4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
