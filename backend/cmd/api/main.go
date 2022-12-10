package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// Define a config struct to hold all the configuration settings for our application.
type config struct {
	port int
	env  string
	db   struct {
		dsn          string
		maxOpenConns int
		maxIdleConns int
		maxIdleTime  string
	}
}

// Define an application struct to hold the dependencies for our HTTP handlers, helpers,
// and middleware.
type application struct {
	Domain   string
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", getIntEnv("PORT"), "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.StringVar(&cfg.db.dsn, "dsn", getStrEnv("DSN_CONN"), "PostgreSQL data source name")

	flag.IntVar(&cfg.db.maxOpenConns, "db-max-open-conns", 25, "PostgreSQL max open connections")
	flag.IntVar(&cfg.db.maxIdleConns, "db-max-idle-conns", 25, "PostgreSQL max idle connections")
	flag.StringVar(&cfg.db.maxIdleTime, "db-max-idle-time", "15m", "PostgreSQL max connection idle time")

	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// Connect to database
	conn, err := openDB(cfg)
	if err != nil {
		infoLog.Fatal(err)
	}
	defer conn.Close()

	infoLog.Printf("database connection pool established")

	app := &application{
		Domain:  "example.com",
		infoLog: infoLog,
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		ErrorLog:     errorLog,
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	infoLog.Printf("Starting %s server on port %s\n", cfg.env, srv.Addr)
	err = srv.ListenAndServe()
	if err != nil {
		errorLog.Fatal(err)
	}
}
