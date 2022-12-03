package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime/debug"
	"strconv"

	"github.com/joho/godotenv"
)

type envelope map[string]any

// writeJSON sends responses: it takes the destination
// http.ResponseWriter, HTTP status code, data to encode to JSON, and a
// header map containing any additional HTTP headers we want to include in the response.
func (app *application) writeJSON(w http.ResponseWriter, status int, data envelope, headers http.Header) error {
	js, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	js = append(js, '\n')

	// At this point, we know that we won't encounter any more errors before writing the
	// response, so it's safe to add any headers that we want to include. We loop
	// through the header map and add each header to the http.ResponseWriter header map.
	for key, value := range headers {
		w.Header()[key] = value
	}

	// Add the "Content-Type: application/json" header, then write the status code and
	// JSON response.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)

	return nil
}

// The serverError helper writes an error message and stack trace to the errorLog
func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Print(2, trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// getStrEnv reads from the environment variables & returns env as a string
func getStrEnv(key string) string {
	envErr := godotenv.Load(".env")
	if envErr != nil {
		log.Fatalf("error loading .env file: %v", envErr)
	}
	val := os.Getenv(key)
	if val == "" {
		panic("error reading environment variable")
	}
	return val
}

// getIntEnv converts env to int
func getIntEnv(key string) int {
	val := getStrEnv(key)
	env, err := strconv.Atoi(val)
	if err != nil {
		panic("error converting environment variable")
	}
	return env
}
