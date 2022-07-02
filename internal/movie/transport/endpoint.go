package transport

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/nrnc/moviesap/internal/entity"
	"github.com/nrnc/moviesap/internal/movie/dto"
)

type Endpoints struct {
	CreateMovie    endpoint.Endpoint
	GetMovies      endpoint.Endpoint
	UpdateMovie    endpoint.Endpoint
	DeleteMovie    endpoint.Endpoint
	GetMovieById   endpoint.Endpoint
	GetMovieByName endpoint.Endpoint
}

func MakeServerEndpoints(s entity.MovieService) Endpoints {
	return Endpoints{
		CreateMovie:  makeCreateMovieEndpoint(s),
		GetMovies:    makeGetMoviesEndpoint(s),
		UpdateMovie:  makeUpdateMovieEndpoint(s),
		GetMovieById: makeGetMovieByIdEndpoint(s),
		DeleteMovie:  makeDeleteMovieEndpoint(s),
	}
}

func makeCreateMovieEndpoint(s entity.MovieService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(dto.CreateMovieRequest)
		id, err := s.CreateMovie(ctx, &req.Movie)
		if err != nil {
			return nil, err
		}
		return dto.CreateMovieResponse{Id: id, Err: nil}, nil
	}
}
func makeGetMoviesEndpoint(s entity.MovieService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		movies, err := s.GetMovies(ctx)
		if err != nil {
			return nil, err
		}
		return dto.GetMoviesResponse{Movies: movies}, nil
	}
}
func makeUpdateMovieEndpoint(s entity.MovieService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(dto.UpdateMovieRequest)
		err = s.UpdateMovie(ctx, &req.Movie)
		if err != nil {
			return dto.UpdateMovieResponse{Success: false, Err: err}, err
		}
		return dto.UpdateMovieResponse{Success: true, Err: nil}, nil
	}
}

func makeDeleteMovieEndpoint(s entity.MovieService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return nil, nil
	}
}

func makeGetMovieByIdEndpoint(s entity.MovieService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(dto.GetMovieByIDRequest)
		movie, err := s.GetMovieByID(ctx, req.ID)
		if err != nil {
			return dto.GetMovieByIDResponse{Movie: entity.Movie{}, Err: err}, err
		}
		return dto.GetMovieByIDResponse{Movie: movie, Err: nil}, nil
	}
}
