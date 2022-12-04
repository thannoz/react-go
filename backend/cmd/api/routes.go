package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/", app.Home)
	router.HandlerFunc(http.MethodGet, "/movies", app.Movies)

	standard := alice.New(app.enableCORS, app.recoverPanic, app.logRequest, secureHeaders)
	return standard.Then(router)
}
