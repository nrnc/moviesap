package entity

import (
	"context"
	"time"
)

type Movie struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	UpdatedAt   time.Time `json:"updated_at"`
	CreatedAt   time.Time `json:"created_at"`
}

type MovieService interface {
	GetMovies(ctx context.Context) ([]Movie, error)
	CreateMovie(ctx context.Context, movie *Movie) (int, error)
	UpdateMovie(ctx context.Context, movie *Movie) error
	DeleteMovie(ctx context.Context, id int64) error
	GetMovieByID(ctx context.Context, id int64) (Movie, error)
}

type MovieRepository interface {
	GetMovies(ctx context.Context) ([]Movie, error)
	CreateMovie(ctx context.Context, movie *Movie) (int, error)
	UpdateMovie(ctx context.Context, movie *Movie) error
	DeleteMovie(ctx context.Context, id int64) error
	GetMovieByID(ctx context.Context, id int64) (Movie, error)
}
