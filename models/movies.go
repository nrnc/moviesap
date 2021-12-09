package models

import (
	"context"
	"database/sql"
	"time"
)

type DBModel struct {
	DB *sql.DB
}

// GetMovieByID fetches the movie by ID
func (m *DBModel) GetMovieByID(id int) (*Movie, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select id,title,description,
				created_at,updated_at from movies where id=$1`
	row := m.DB.QueryRowContext(ctx, query, id)
	var movie Movie
	err := row.Scan(
		&movie.ID,
		&movie.Title,
		&movie.Description,
		&movie.CreatedAt,
		&movie.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &movie, nil
}

// InsertMovie inserts a movie into the database
func (m *DBModel) InsertMovie(h Movie) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var newId int
	stmt := `insert into movies (title,description,created_at,updated_at) values($1,$2,$3,$4)
				returning id`
	err := m.DB.QueryRowContext(ctx, stmt,
		h.Title,
		h.Description,
		time.Now(),
		time.Now(),
	).Scan(&newId)
	if err != nil {
		return 0, err
	}
	return newId, nil
}

func (m *DBModel) UpdateMovie(h Movie) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	query := `
		update movies set title=$1,description=$2,created_at=$3,updated_at=$4 where id = $5
	`

	_, err := m.DB.ExecContext(ctx, query, h.Title, h.Description, time.Now(), time.Now(), h.ID)

	if err != nil {
		return err
	}
	return nil

}

func (m *DBModel) DeleteMovieByID(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	query := `
		delete from movies where id=$1
	`

	_, err := m.DB.ExecContext(ctx, query, id)

	if err != nil {
		return err
	}
	return nil
}
