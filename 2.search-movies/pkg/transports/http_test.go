package transports_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"search-movies/pkg/model"

	"github.com/stretchr/testify/assert"
)

func doRequest(req *http.Request) *httptest.ResponseRecorder {
	responseRecorder := httptest.NewRecorder()
	httpServerTest.ServeHTTP(responseRecorder, req)
	return responseRecorder
}

func TestSearchMoviesHandler(t *testing.T) {
	expectedResult := &model.SearchResponse{
		Movies: []model.Movie{
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

	searchRequest := &model.SearchRequest{
		Pagination: 2,
		SearchWord: "batman",
	}
	reqByte, _ := json.Marshal(searchRequest)
	req, err := http.NewRequest("GET", "/search", bytes.NewBuffer(reqByte))
	if err != nil {
		t.Fatal(err)
	}

	responseRecorder := doRequest(req)

	assert.Equal(t, http.StatusOK, responseRecorder.Code,
		"response status code should be %v", http.StatusOK)

	decoder := json.NewDecoder(responseRecorder.Body)
	var searchResponse model.SearchResponse
	_ = decoder.Decode(&searchResponse)
	assert.Equal(t, expectedResult.Movies, searchResponse.Movies,
		"movies should be %v", expectedResult.Movies)
	assert.Equal(t, expectedResult.Err, searchResponse.Err,
		"error should be %v", expectedResult.Err)
}

func TestHealthCheckHandler(t *testing.T) {
	healthCheckRequest := &model.HealthCheckRequest{}
	reqByte, _ := json.Marshal(healthCheckRequest)
	req, err := http.NewRequest("GET", "/healthz", bytes.NewBuffer(reqByte))
	if err != nil {
		t.Fatal(err)
	}

	responseRecorder := doRequest(req)

	assert.Equal(t, http.StatusOK, responseRecorder.Code,
		"response status code should be %v", http.StatusOK)
	decoder := json.NewDecoder(responseRecorder.Body)
	var healthResponse model.HealthCheckResponse
	_ = decoder.Decode(&healthResponse)
	assert.Equal(t, int64(http.StatusOK), healthResponse.Code,
		"response code should be %v", int64(http.StatusOK))
}
