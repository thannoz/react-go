package main

import (
	"fmt"
	"net/http"
)

// test method
func (app *application) Home(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Hello, world from %s\n", app.Domain)
}
