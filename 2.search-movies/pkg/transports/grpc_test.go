package transports_test

import (
	"context"
	"net/http"
	"testing"

	"search-movies/api/v1/searchMovies"

	"github.com/stretchr/testify/assert"
)

func TestGrpcServer_Search(t *testing.T) {
	expectedResult := &searchMovies.SearchResponse{
		Movies: []*searchMovies.Movie{
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
		},
		Err: "",
	}

	searchRequest := &searchMovies.SearchRequest{
		Pagination: 2,
		SearchWord: "batman",
	}

	resp, err := gRpcServerTest.Search(context.Background(), searchRequest)

	assert.Empty(t, err, "error should be empty")
	assert.Equal(t, expectedResult.Movies, resp.Movies,
		"movies should be %v", expectedResult.Movies)
	assert.Equal(t, expectedResult.Err, resp.Err,
		"error should be %v", expectedResult.Err)
}

func TestGrpcServer_HealthCheck(t *testing.T) {
	expectedResult := &searchMovies.HealthCheckResponse{
		Code: int64(http.StatusOK),
		Err:  "",
	}

	healthCheckRequest := &searchMovies.HealthCheckRequest{}

	resp, err := gRpcServerTest.HealthCheck(context.Background(), healthCheckRequest)

	assert.Empty(t, err, "error should be empty")
	assert.Equal(t, expectedResult.Code, resp.Code,
		"code should be %v", expectedResult.Code)
	assert.Equal(t, expectedResult.Err, resp.Err,
		"error should be %v", expectedResult.Err)
}
