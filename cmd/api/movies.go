package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/van9md/greenlight/internal/data"
)

func (app *application) createMoviesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new movie")
}

func (app *application) showMoviesHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFountResponse(w, r)
		return
	}
	movie := data.Movie{
		ID:        id,
		CreatedAt: time.Now(),
		Title:     "Deadpool",
		Runtime:   169,
		Genres:    []string{"comedy", "action"},
		Version:   1,
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"movie": movie}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
