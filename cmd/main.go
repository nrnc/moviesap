package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/lib/pq"
	"github.com/nrnc/moviesap/internal/entity"
	"github.com/nrnc/moviesap/internal/movie/moviesvc"
	"github.com/nrnc/moviesap/internal/movie/repository"
	"github.com/nrnc/moviesap/internal/movie/transport"
	httptransport "github.com/nrnc/moviesap/internal/movie/transport/http"
)

type config struct {
	port int
	env  string
	db   struct {
		dsn string
	}
}

func main() {
	var cfg config
	flag.IntVar(&cfg.port, "port", 4000, "Server port to listen on")
	flag.StringVar(&cfg.env, "env", "development", "Aplication Environment(production|development)")
	flag.StringVar(&cfg.db.dsn, "dsn", "postgres://unbxd:password@localhost:5432/movies?sslmode=disable", "postgres connection string")
	flag.Parse()
	fmt.Println(cfg.db.dsn)
	db, err := openDB(cfg)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()
	var s entity.MovieService
	{
		repo := repository.New(db)
		s = moviesvc.New(repo, time.Second)
	}
	var endpoints transport.Endpoints
	{
		endpoints = transport.MakeServerEndpoints(s)
	}
	var h http.Handler
	{
		h = httptransport.NewService(endpoints)
	}

	err = http.ListenAndServe(":4004", h)
	if err != nil {
		log.Fatal(err)
	}
}

func openDB(config config) (*sql.DB, error) {
	db, err := sql.Open("postgres", config.db.dsn)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}
