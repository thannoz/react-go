package main

import (
	"fmt"
	"net/http"
)

func (app *application) Home(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Hello, world from %s\n", app.Domain)
}

func (app *application) Movies(w http.ResponseWriter, r *http.Request) {
	movies, err := app.DB.AllMovies()
	if err != nil {
		app.serverError(w, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"movies": movies}, nil)
	if err != nil {
		app.errorLog.Print(err)
		app.serverError(w, err)
	}
}
