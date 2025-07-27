package main

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/Infamous003/snippetbox/internal/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// files := []string{
	// 	"./ui/html/base.html",
	// 	"./ui/html/partials/nav.html",
	// 	"./ui/html/pages/home.html",
	// }

	snippets, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}

	for _, snippet := range snippets {
		fmt.Fprintf(w, "%+v\n", snippet)
	}

	// temps, err := template.ParseFiles(files...)

	// if err != nil {
	// 	app.serverError(w, err)
	// 	return
	// }

	// err = temps.ExecuteTemplate(w, "base", nil)

	// if err != nil {
	// 	app.serverError(w, err)
	// }
}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	strId := r.URL.Query().Get("id")
	id, err := strconv.Atoi(strId)

	if err != nil {
		app.errorLog.Println(err)
		return
	}

	snippet, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
	}

	files := []string{
		"./ui/html/base.html",
		"./ui/html/partials/nav.html",
		"./ui/html/pages/view.html",
	}

	// Here we are parsing the foles
	tmpls, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}
	data := &templateData{
		Snippet: snippet,
	}
	err = tmpls.ExecuteTemplate(w, "base", data)
	if err != nil {
		app.serverError(w, err)
	}
}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		// http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	title := "First autumn morning"
	content := "First autumn morning\nthe mirror I stare into\nshows my father''s face.\n\n– Murakami Kijo"
	expires := 8

	id, err := app.snippets.Insert(title, content, expires)

	if err != nil {
		app.serverError(w, err)
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/snippet/view?id=%d", id), http.StatusSeeOther)
}
