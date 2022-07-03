package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/nrnc/moviesap/internal/entity"
)

type movieRepository struct {
	Conn *sql.DB
}

func New(conn *sql.DB) *movieRepository {
	return &movieRepository{
		Conn: conn,
	}
}

func (m *movieRepository) CreateMovie(ctx context.Context, movie *entity.Movie) (id int, err error) {
	stmt := `insert into movies (title,description,created_at,updated_at) values($1,$2,$3,$4)
				returning id`
	err = m.Conn.QueryRowContext(ctx, stmt,
		movie.Name,
		movie.Description,
		time.Now(),
		time.Now(),
	).Scan(&id)
	if err != nil {
		return id, err
	}
	return
}
func (m *movieRepository) GetMovies(ctx context.Context) ([]entity.Movie, error) {
	query := `select id,title,description,created_at,updated_at from movies `

	rows, err := m.Conn.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var movies []entity.Movie

	for rows.Next() {
		var movie entity.Movie
		err := rows.Scan(
			&movie.ID,
			&movie.Name,
			&movie.Description,
			&movie.CreatedAt,
			&movie.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		movies = append(movies, movie)
	}

	return movies, nil
}
func (m *movieRepository) UpdateMovie(ctx context.Context, movie *entity.Movie) error {
	query := `
		update movies set title=$1,description=$2,updated_at=$3 where id = $4
	`

	_, err := m.Conn.ExecContext(ctx, query, movie.Name, movie.Description, time.Now(), movie.ID)

	if err != nil {
		return err
	}

	return nil
}
func (m *movieRepository) DeleteMovie(ctx context.Context, id int64) error {
	query := `
		delete from movies where id=$1
	`

	_, err := m.Conn.ExecContext(ctx, query, id)

	if err != nil {
		return err
	}
	return nil
}
func (m *movieRepository) GetMovieByID(ctx context.Context, id int64) (entity.Movie, error) {
	query := `select id,title,description,
				created_at,updated_at from movies where id=$1`
	row := m.Conn.QueryRowContext(ctx, query, id)
	var movie entity.Movie
	err := row.Scan(
		&movie.ID,
		&movie.Name,
		&movie.Description,
		&movie.CreatedAt,
		&movie.UpdatedAt,
	)
	if err != nil {
		return movie, entity.ErrNotFound
	}

	return movie, nil
}
