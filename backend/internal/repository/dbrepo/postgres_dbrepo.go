package dbrepo

import (
	"backend/internal/models"
	"context"
	"database/sql"
	"time"
)

type PostgresDBRepo struct {
	DB *sql.DB
}

const dbTimeout = time.Second * 3

func (m *PostgresDBRepo) AllMovies() ([]*models.Movie, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
	SELECT
		id, title, release_date, runtime,
		mpaa_rating, description, coalesce(image, ''),
		created_at, updated_at
	FROM 
		movies
	ORDER BY
		title
	`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	movies := []*models.Movie{}
	for rows.Next() {
		movie := &models.Movie{}
		if err := rows.Scan(
			&movie.ID,
			&movie.Title,
			&movie.ReleaseDate,
			&movie.Runtime,
			&movie.MPAARating,
			&movie.Description,
			&movie.Image,
			&movie.CreatedAt,
			&movie.UpdatedAt,
		); err != nil {
			return nil, err
		}
		movies = append(movies, movie)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return movies, nil
}
