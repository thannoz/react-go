package main

import (
	"context"
	"database/sql"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func openDB(cfg config) (*sql.DB, error) {
	db, err := sql.Open("pgx", cfg.db.dsn)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(cfg.db.maxOpenConns)

	db.SetMaxIdleConns(cfg.db.maxIdleConns)

	// Convert idle timeout duration string to time.Duration
	duration, err := time.ParseDuration(cfg.db.maxIdleTime)
	if err != nil {
		return nil, err
	}

	db.SetConnMaxIdleTime(duration)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Use PingContext() to establish a new connection to the database
	if err = db.PingContext(ctx); err != nil {
		return nil, err
	}
	return db, nil
}

// func (app *application) connectToDB() (*sql.DB, error) {
// 	conn, err := openDB(app.cfg)
// 	if err != nil {
// 		return nil, err
// 	}

// 	app.infoLog.Println("Connecting to PostgreSQL")

// 	return conn, nil
// }
