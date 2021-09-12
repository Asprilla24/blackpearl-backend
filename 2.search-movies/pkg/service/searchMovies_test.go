package service_test

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"testing"

	"search-movies/pkg/config"
	"search-movies/pkg/dao"
	"search-movies/pkg/dao/postgres"
	"search-movies/pkg/model"
	"search-movies/pkg/service"

	"github.com/caarlos0/env"
	"github.com/go-kit/kit/log"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

var (
	dbConnTest      *gorm.DB
	confTest        *config.Config
	dbTest          dao.DB
	httpServiceTest service.Service
)

// nolint : dupl
func TestMain(m *testing.M) {
	var logger log.Logger

	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "search-movies", log.DefaultTimestampUTC)

	confTest = &config.Config{}
	flag.Usage = func() {
		flag.CommandLine.SetOutput(os.Stdout)
		for _, val := range confTest.HelpDocs() {
			fmt.Println(val)
		}
		fmt.Println("")
		flag.PrintDefaults()
	}
	flag.Parse()

	err := env.Parse(confTest)
	if err != nil {
		logger.Log("ENV", "Parse", "err", err)
		return
	}

	dbConnTest, err = dao.NewPostgres("postgres", confTest)
	if err != nil {
		logger.Log("DAO", "NewPostgres", "err", err)
		return
	}
	defer dbConnTest.Close() // nolint : errcheck, used in defer

	dbTest = postgres.NewDB(dbConnTest)
	dbTest.MigrateDB(&model.Log{})

	httpServiceTest = service.NewService(dbTest, *confTest)

	os.Exit(m.Run())
}

func TestSearchMoviesHandler(t *testing.T) {
	expectedResult := []model.Movie{
		{
			Title:  "Batman: The Killing Joke",
			Year:   "2016",
			ImdbID: "tt4853102",
			Type:   "movie",
			Poster: "https://m.media-amazon.com/images/M/MV5BMTdjZTliODYtNWExMi00NjQ1LWIzN2MtN2Q5NTg5NTk3NzliL2ltYWdlXkEyXkFqcGdeQXVyNTAyODkwOQ@@._V1_SX300.jpg",
		},
		{
			Title:  "Batman: The Dark Knight Returns, Part 2",
			Year:   "2013",
			ImdbID: "tt2166834",
			Type:   "movie",
			Poster: "https://m.media-amazon.com/images/M/MV5BYTEzMmE0ZDYtYWNmYi00ZWM4LWJjOTUtYTE0ZmQyYWM3ZjA0XkEyXkFqcGdeQXVyNTA4NzY1MzY@._V1_SX300.jpg",
		},
		{
			Title:  "Batman: Mask of the Phantasm",
			Year:   "1993",
			ImdbID: "tt0106364",
			Type:   "movie",
			Poster: "https://m.media-amazon.com/images/M/MV5BYTRiMWM3MGItNjAxZC00M2E3LThhODgtM2QwOGNmZGU4OWZhXkEyXkFqcGdeQXVyNjExODE1MDc@._V1_SX300.jpg",
		},
		{
			Title:  "Batman: Year One",
			Year:   "2011",
			ImdbID: "tt1672723",
			Type:   "movie",
			Poster: "https://m.media-amazon.com/images/M/MV5BNTJjMmVkZjctNjNjMS00ZmI2LTlmYWEtOWNiYmQxYjY0YWVhXkEyXkFqcGdeQXVyNTAyODkwOQ@@._V1_SX300.jpg",
		},
		{
			Title:  "Batman: Assault on Arkham",
			Year:   "2014",
			ImdbID: "tt3139086",
			Type:   "movie",
			Poster: "https://m.media-amazon.com/images/M/MV5BZDU1ZGRiY2YtYmZjMi00ZDQwLWJjMWMtNzUwNDMwYjQ4ZTVhXkEyXkFqcGdeQXVyNTAyODkwOQ@@._V1_SX300.jpg",
		},
		{
			Title:  "Batman: The Movie",
			Year:   "1966",
			ImdbID: "tt0060153",
			Type:   "movie",
			Poster: "https://m.media-amazon.com/images/M/MV5BMmM1OGIzM2UtNThhZS00ZGNlLWI4NzEtZjlhOTNhNmYxZGQ0XkEyXkFqcGdeQXVyNTkxMzEwMzU@._V1_SX300.jpg",
		},
		{
			Title:  "Batman: Gotham Knight",
			Year:   "2008",
			ImdbID: "tt1117563",
			Type:   "movie",
			Poster: "https://m.media-amazon.com/images/M/MV5BM2I0YTFjOTUtMWYzNC00ZTgyLTk2NWEtMmE3N2VlYjEwN2JlXkEyXkFqcGdeQXVyNTAyODkwOQ@@._V1_SX300.jpg",
		},
		{
			Title:  "Batman: Arkham City",
			Year:   "2011",
			ImdbID: "tt1568322",
			Type:   "game",
			Poster: "https://m.media-amazon.com/images/M/MV5BZDE2ZDFhMDAtMDAzZC00ZmY3LThlMTItMGFjMzRlYzExOGE1XkEyXkFqcGdeQXVyNTAyODkwOQ@@._V1_SX300.jpg",
		},
		{
			Title:  "Batman Beyond",
			Year:   "1999–2001",
			ImdbID: "tt0147746",
			Type:   "series",
			Poster: "https://m.media-amazon.com/images/M/MV5BYTBiZjFlZDQtZjc1MS00YzllLWE5ZTQtMmM5OTkyNjZjMWI3XkEyXkFqcGdeQXVyMTA1OTEwNjE@._V1_SX300.jpg",
		},
		{
			Title:  "Son of Batman",
			Year:   "2014",
			ImdbID: "tt3139072",
			Type:   "movie",
			Poster: "https://m.media-amazon.com/images/M/MV5BYjdkZWFhNzctYmNhNy00NGM5LTg0Y2YtZWM4NmU2MWQ3ODVkXkEyXkFqcGdeQXVyNTA0OTU0OTQ@._V1_SX300.jpg",
		},
	}

	searchRequest := &model.SearchRequest{
		Pagination: 2,
		SearchWord: "batman",
	}

	resp, err := httpServiceTest.Search(context.Background(), searchRequest.Pagination, searchRequest.SearchWord)

	assert.Empty(t, err,
		"error should be empty")
	assert.Equal(t, expectedResult, resp,
		"movies should be %v", expectedResult)
}

func TestHealthCheckHandler(t *testing.T) {
	resp, err := httpServiceTest.HealthCheck(context.Background())

	assert.Empty(t, err,
		"error should be empty")
	assert.Equal(t, int64(http.StatusOK), resp,
		"movies should be %v", int64(http.StatusOK))
}
