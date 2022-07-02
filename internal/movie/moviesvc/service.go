package moviesvc

import (
	"context"
	"time"

	"github.com/nrnc/moviesap/internal/entity"
)

type moviesService struct {
	mrepo   entity.MovieRepository
	timeout time.Duration
}

func (m *moviesService) CreateMovie(ctx context.Context, movie *entity.Movie) (id int, err error) {
	ctx, cancel := context.WithTimeout(ctx, m.timeout)
	defer cancel()

	id, err = m.mrepo.CreateMovie(ctx, movie)
	return
}
func (m *moviesService) GetMovies(ctx context.Context) ([]entity.Movie, error) {
	ctx, cancel := context.WithTimeout(ctx, m.timeout)
	defer cancel()
	return m.mrepo.GetMovies(ctx)
}
func (m *moviesService) UpdateMovie(ctx context.Context, movie *entity.Movie) error {
	ctx, cancel := context.WithTimeout(ctx, m.timeout)
	defer cancel()

	err := m.mrepo.UpdateMovie(ctx, movie)
	return err
}
func (m *moviesService) DeleteMovie(ctx context.Context, id int64) error {
	ctx, cancel := context.WithTimeout(ctx, m.timeout)
	defer cancel()

	return m.mrepo.DeleteMovie(ctx, id)
}
func (m *moviesService) GetMovieByID(ctx context.Context, id int64) (entity.Movie, error) {
	ctx, cancel := context.WithTimeout(ctx, m.timeout)
	defer cancel()
	return m.mrepo.GetMovieByID(ctx, id)
}
func New(mrepo entity.MovieRepository, timeout time.Duration) entity.MovieService {
	return &moviesService{
		mrepo:   mrepo,
		timeout: timeout,
	}
}
