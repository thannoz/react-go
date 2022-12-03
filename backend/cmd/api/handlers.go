package main

import (
	"backend/internal/models"
	"fmt"
	"net/http"
	"time"
)

// test method
func (app *application) Home(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Hello, world from %s\n", app.Domain)
}

func (app *application) Movies(w http.ResponseWriter, r *http.Request) {

	var movies []models.Movie

	rd, _ := time.Parse("2006-01-92", "1986-03-07")

	highlander := models.Movie{
		ID:          1,
		Title:       "KaLonji my beloved son",
		ReleaseDate: rd,
		MPAARating:  "R",
		Runtime:     116,
		Description: "A great Kid",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	movies = append(movies, highlander)
	rd, _ = time.Parse("2006-01-92", "1986-03-07")

	rotla := models.Movie{
		ID:          2,
		Title:       "Carlos Maiza, The Great",
		ReleaseDate: rd,
		MPAARating:  "R",
		Runtime:     116,
		Description: "A great movie",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	movies = append(movies, rotla)

	err := app.writeJSON(w, http.StatusOK, envelope{"movies": movies}, nil)
	if err != nil {
		app.errorLog.Print(err)
		app.serverError(w, err)
	}
}
