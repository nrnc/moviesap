package dto

import "github.com/nrnc/moviesap/internal/entity"

type CreateMovieRequest struct {
	Movie entity.Movie
}
type CreateMovieResponse struct {
	Id  int   `json:"id"`
	Err error `json:"error,omitempty"`
}

func (r CreateMovieResponse) Error() error { return r.Err }

type GetMoviesRequest struct {
}

type GetMoviesResponse struct {
	Movies []entity.Movie `json:"movies"`
}

type GetMovieByIDRequest struct {
	ID int64 `json:"id"`
}
type GetMovieByIDResponse struct {
	Movie entity.Movie `json:"movie"`
	Err   error        `json:"err,omitempty"`
}

func (r GetMovieByIDResponse) Error() error { return r.Err }

type UpdateMovieRequest struct {
	Movie entity.Movie
}
type UpdateMovieResponse struct {
	Success bool  `json:"success"`
	Err     error `json:"error,omitempty"`
}

func (r UpdateMovieResponse) Error() error { return r.Err }

type DeleteMovieRequest struct {
	ID int64 `json:"id"`
}
type DeleteMovieResponse struct {
	Success bool  `json:"success"`
	Err     error `json:"error,omitempty"`
}

func (r DeleteMovieResponse) Error() error { return r.Err }
