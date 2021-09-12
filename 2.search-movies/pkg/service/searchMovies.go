package service

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"search-movies/pkg/config"
	"search-movies/pkg/dao"
	"search-movies/pkg/model"
	"search-movies/pkg/util"

	"github.com/go-kit/kit/log"
)

type searchMoviesService struct {
	DB     dao.DB
	Config config.Config
}

func NewService(db dao.DB, config config.Config) Service {
	return &searchMoviesService{DB: db, Config: config}
}

func (service *searchMoviesService) Search(_ context.Context, pagination int64, searchWord string) ([]model.Movie, error) {
	result, err := service.searchMovies(pagination, searchWord)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (service *searchMoviesService) HealthCheck(_ context.Context) (int64, error) {
	if err := service.DB.Health(); err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}

func (service *searchMoviesService) searchMovies(pagination int64, searchWord string) ([]model.Movie, error) {
	getSearchMoviesURL := fmt.Sprintf("http://www.omdbapi.com/?apikey=%s&s=%s&page=%d", service.Config.OMDBKey, searchWord, pagination)

	newReq, err := http.NewRequest(http.MethodGet, getSearchMoviesURL, nil)
	if err != nil {
		return nil, err
	}

	body, status, err := util.DoHTTPRequest(newReq)
	if err != nil {
		return nil, err
	}

	err = service.createLog(getSearchMoviesURL, status)
	if err != nil {
		logger.Log("createLog", err.Error())
	}

	if status != http.StatusOK {
		return nil, fmt.Errorf("got status: %d, when getting search movies", status)
	}

	var respBody model.SearchIMDBResponse
	err = json.Unmarshal(body, &respBody)
	if err != nil {
		return nil, err
	}

	return respBody.Search, nil
}

func (service *searchMoviesService) createLog(url string, respStatus int) error {
	log := model.Log{
		URL:            url,
		ResponseStatus: respStatus,
	}
	return service.DB.InsertLog(log)
}

var logger log.Logger

func init() {
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "service", log.DefaultTimestampUTC)
}
