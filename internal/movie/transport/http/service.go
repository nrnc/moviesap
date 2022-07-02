package http

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/nrnc/moviesap/internal/entity"
	"github.com/nrnc/moviesap/internal/movie/dto"
	svctransport "github.com/nrnc/moviesap/internal/movie/transport"
)

var (
	// ErrBadRouting is returned when an expected path variable is missing.
	ErrBadRouting = errors.New("inconsistent mapping between route and handler (programmer error)")
)

func NewService(enpoints svctransport.Endpoints) http.Handler {
	r := mux.NewRouter()

	r.Methods("POST").Path("/movies").Handler(kithttp.NewServer(
		enpoints.CreateMovie,
		decodeCreateMovieRequest,
		encodeResponse,
	))
	r.Methods("GET").Path("/movies/{id}").Handler(kithttp.NewServer(
		enpoints.GetMovieById,
		decodeGetMovieByIdRequest,
		encodeResponse,
	))
	r.Methods("GET").Path("/movies").Handler(kithttp.NewServer(
		enpoints.GetMovies,
		decodeGetMoviesRequest,
		encodeResponse,
	))
	return r
}

func decodeCreateMovieRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req dto.CreateMovieRequest
	if e := json.NewDecoder(r.Body).Decode(&req.Movie); e != nil {
		return nil, e
	}
	return req, nil
}
func decodeGetMoviesRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req dto.GetMoviesRequest
	return req, nil
}

func decodeGetMovieByIdRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req dto.GetMovieByIDRequest
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	req.ID = int64(id)
	return req, err
}

type Errorer interface {
	Error() error
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(Errorer); ok && e.Error() != nil { // Not a Go kit transport error, but a business-logic error.
		// Provide those as HTTP errors.
		encodeError(ctx, e.Error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}
func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("encodeError with nil error")
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(codeFrom(err))
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}
func codeFrom(err error) int {
	switch err {
	case entity.ErrNotFound:
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}
