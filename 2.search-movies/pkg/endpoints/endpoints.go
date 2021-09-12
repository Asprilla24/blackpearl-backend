package endpoints

import (
	"context"
	"errors"
	"os"

	"search-movies/pkg/model"
	"search-movies/pkg/service"
	"search-movies/pkg/util"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
)

type Set struct {
	SearchEndpoint      endpoint.Endpoint
	HealthCheckEndpoint endpoint.Endpoint
}

func NewEndpointSet(svc service.Service) Set {
	return Set{
		SearchEndpoint:      MakeSearchEndpoint(svc),
		HealthCheckEndpoint: MakeHealthCheckEndpoint(svc),
	}
}

func MakeSearchEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(model.SearchRequest)
		if !ok {
			logger.Log("MakeSearchEndpoint", "request error")
			return nil, util.ErrBadRequest
		}

		movies, err := svc.Search(ctx, req.Pagination, req.SearchWord)
		if err != nil {
			logger.Log("MakeSearchEndpoint error:", err.Error())
			return model.SearchResponse{Movies: movies, Err: err.Error()}, nil
		}
		return model.SearchResponse{Movies: movies, Err: ""}, nil
	}
}

func MakeHealthCheckEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_, ok := request.(model.HealthCheckRequest)
		if !ok {
			logger.Log("MakeHealthCheckEndpoint", "request error")
			return nil, util.ErrBadRequest
		}

		code, err := svc.HealthCheck(ctx)
		if err != nil {
			logger.Log("MakeHealthCheckEndpoint, error:", err.Error())
			return model.HealthCheckResponse{Code: code, Err: err.Error()}, nil
		}
		return model.HealthCheckResponse{Code: code, Err: ""}, nil
	}
}

func (s *Set) Search(ctx context.Context, pagination int64, searchWord string) ([]model.Movie, error) {
	resp, err := s.SearchEndpoint(ctx, model.SearchRequest{Pagination: pagination, SearchWord: searchWord})
	if err != nil {
		logger.Log("Search, error:", err.Error())
		return []model.Movie{}, err
	}
	searchResp, ok := resp.(model.SearchResponse)
	if !ok {
		logger.Log("Search", "response error")
		return nil, errors.New("search response error")
	}
	if searchResp.Err != "" {
		return []model.Movie{}, errors.New(searchResp.Err)
	}
	return searchResp.Movies, nil
}

func (s *Set) HealthCheck(ctx context.Context) (int64, error) {
	resp, err := s.HealthCheckEndpoint(ctx, model.HealthCheckRequest{})
	if err != nil {
		logger.Log("HealthCheck, error:", err.Error())
		return 0, err
	}
	healthCheckResp, ok := resp.(model.HealthCheckResponse)
	if !ok {
		logger.Log("HealthCheck", "response error")
		return 0, errors.New("healthcheck response error")
	}

	if healthCheckResp.Err != "" {
		return healthCheckResp.Code, errors.New(healthCheckResp.Err)
	}
	return healthCheckResp.Code, nil
}

var logger log.Logger

func init() {
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "endpoints", log.DefaultTimestampUTC)
}
