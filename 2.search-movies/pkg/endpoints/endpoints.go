package endpoints

import (
	"context"
	"errors"
	"search-movies/pkg/model"
	"search-movies/pkg/service"
	"search-movies/pkg/util"

	"github.com/go-kit/kit/endpoint"
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
			return nil, util.ErrBadRequest
		}

		movies, err := svc.Search(ctx, req.Pagination, req.SearchWord)
		if err != nil {
			return model.SearchResponse{Movies: movies, Err: err.Error()}, nil
		}
		return model.SearchResponse{Movies: movies, Err: ""}, nil
	}
}

func MakeHealthCheckEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_, ok := request.(model.HealthCheckRequest)
		if !ok {
			return nil, util.ErrBadRequest
		}

		code, err := svc.HealthCheck(ctx)
		if err != nil {
			return model.HealthCheckResponse{Code: code, Err: err.Error()}, nil
		}
		return model.HealthCheckResponse{Code: code, Err: ""}, nil
	}
}

func (s *Set) Search(ctx context.Context, pagination int64, searchWord string) ([]model.Movie, error) {
	resp, err := s.SearchEndpoint(ctx, model.SearchRequest{Pagination: pagination, SearchWord: searchWord})
	if err != nil {
		return []model.Movie{}, err
	}
	searchResp := resp.(model.SearchResponse)
	if searchResp.Err != "" {
		return []model.Movie{}, errors.New(searchResp.Err)
	}
	return searchResp.Movies, nil
}

func (s *Set) HealthCheck(ctx context.Context) (int64, error) {
	resp, err := s.HealthCheckEndpoint(ctx, model.HealthCheckRequest{})
	healthCheckResp := resp.(model.HealthCheckResponse)
	if err != nil {
		return healthCheckResp.Code, err
	}
	if healthCheckResp.Err != "" {
		return healthCheckResp.Code, errors.New(healthCheckResp.Err)
	}
	return healthCheckResp.Code, nil
}
