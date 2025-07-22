package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

func (app *application) serverError(w http.ResponseWriter, err error) {
	//err.Error gives us the error text from the error we passed in
	//debug.Stack gives us the details of the funcs that were running, where the error happened, how the program got there etc
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())

	// app.errorLog.Println(trace)
	//This helps us go back one step in the stack trace. We want the error location to be one step above. the line above will make it so that when an error occurs, the line that it points to would be helpers line 14, but we want it to go to the handler function itself, where we are referencing the base.html file
	app.errorLog.Output(2, trace)

	// We could hardcode the status text ourselves but vuilding it like this is nice
	// smae for the status code, its similar to status.HTTP_404 from fastapi
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}
