package main

import (
	"context"
	"database/sql"
	"db-access/models"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	net_http "net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/lib/pq"
	"github.com/unbxd/go-base/kit/transport/http"
)

type config struct {
	port int
	env  string
	db   struct {
		dsn string
	}
}
type AppStatus struct {
	Status      string `json:"status"`
	Environment string `json:"environment"`
}

type application struct {
	models models.Models
}

func main() {
	var cfg config
	flag.IntVar(&cfg.port, "port", 4000, "Server port to listen on")
	flag.StringVar(&cfg.env, "env", "development", "Aplication Environment(production|development)")
	flag.StringVar(&cfg.db.dsn, "dsn", "postgres://unbxd@localhost/movies?sslmode=disable", "postgres connection string")
	flag.Parse()
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	db, err := openDB(cfg)
	if err != nil {
		logger.Fatal(err)
	}
	defer db.Close()
	movies := models.NewModels(db)
	server, err := http.NewTransport(
		"0.0.0.0",
		"4444",
	)
	if err != nil {
		logger.Fatal(err)
	}
	server.Get("/health-check", func(ctx context.Context, req *net_http.Request) (*net_http.Response, error) {
		return http.NewResponse(
			req,
			http.ResponseWithBytes([]byte("Health Check")),
		), err
	})

	// to read this business object from the network object, i.e. http.Request
	// Decoder is used
	// Decoder is defined as `func(context.Context, *net_http.Request) (interface{}, error)`
	// where in the function reads the net_http.Request and translates it into B1
	decoderFunc := func(
		_ context.Context, req *net_http.Request,
	) (interface{}, error) {
		// here we read req.Body into the object Employee
		var movie models.Movie

		err := json.NewDecoder(req.Body).Decode(&movie)
		if err != nil {
			return nil, fmt.Errorf("error in decoding: %s", err.Error())
		}
		// once decodes succeeds, the Business Object B1 is read from
		// request.Body and decoded in emp
		return movie, nil
	}

	// The final step in this whole flow would be then to convert Manager to
	// corresponding JSON. This step is called Encoding, and is done by http.Encoder
	encodingFunc := func(_ context.Context, rw net_http.ResponseWriter, b2 interface{}) error {
		// b2 is the business object mentioned above.
		// b2 needs to be type casted
		manager := b2.(models.Movie)

		// http packages comes with highly efficient default encoder
		// which makes sure that the data is properly copied onto the response writer
		// We can leverage that implementation or we can write the entire functionality
		// from scratch.

		bt, err := json.Marshal(manager)
		if err != nil {
			return err
		}

		rw.WriteHeader(net_http.StatusOK)
		rw.Write(bt)
		return nil
	}

	server.POST("/movies", func(cx context.Context, req interface{}) (interface{}, error) {
		movies.DB.InsertMovie(req.(models.Movie))
		return req.(models.Movie), err
	}, http.HandlerWithDecoder(decoderFunc),
		http.HandlerWithEncoder(encodingFunc),
	)

	server.Get("/movies/:id", func(ctx context.Context, req *net_http.Request) (*net_http.Response, error) {
		params := http.Parameters(req)
		id, err := strconv.Atoi(params.ByName("id"))
		if err != nil {
			return nil, err
		}
		movie, _ := movies.DB.GetMovieByID(id)
		resp, _ := json.Marshal(movie)

		return http.NewResponse(req, http.ResponseWithBytes(resp)), nil

	})
	server.Put("/movies/:id", func(ctx context.Context, req *net_http.Request) (*net_http.Response, error) {
		params := http.Parameters(req)
		id, err := strconv.Atoi(params.ByName("id"))
		if err != nil {
			return nil, err
		}
		fmt.Println(id)
		movie, err := movies.DB.GetMovieByID(id)
		fmt.Println(movie)
		if err != nil {
			return nil, err
		}
		err = json.NewDecoder(req.Body).Decode(&movie)
		fmt.Println(movie)
		if err != nil {
			return nil, err
		}
		err = movies.DB.UpdateMovie(*movie)
		resp, _ := json.Marshal(*movie)
		if err != nil {
			return nil, err
		}
		return http.NewResponse(req, http.ResponseWithBytes(resp)), nil
	})
	server.Delete("/movies/:id", func(ctx context.Context, req *net_http.Request) (*net_http.Response, error) {
		params := http.Parameters(req)
		id, err := strconv.Atoi(params.ByName("id"))
		if err != nil {
			return nil, err
		}
		movies.DB.DeleteMovieByID(id)
		return http.NewResponse(req, http.ResponseWithBytes([]byte("Ok"))), nil
	})
	server.Open()

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
