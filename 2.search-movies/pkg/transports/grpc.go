package transports

import (
	"context"
	"errors"

	"search-movies/api/v1/searchMovies"
	"search-movies/pkg/endpoints"
	"search-movies/pkg/model"

	grpctransport "github.com/go-kit/kit/transport/grpc"
)

type grpcServer struct {
	search     grpctransport.Handler
	heathCheck grpctransport.Handler
}

func NewGRPCServer(ep endpoints.Set) searchMovies.SearchMoviesServer {
	return &grpcServer{
		search: grpctransport.NewServer(
			ep.SearchEndpoint,
			decodeGRPCSearchRequest,
			decodeGRPCSearchResponse,
		),
		heathCheck: grpctransport.NewServer(
			ep.HealthCheckEndpoint,
			decodeGRPCHealthCheckRequest,
			decodeGRPCHealthCheckResponse,
		),
	}
}

func (g *grpcServer) Search(ctx context.Context, r *searchMovies.SearchRequest) (*searchMovies.SearchResponse, error) {
	_, rep, err := g.search.ServeGRPC(ctx, r)
	if err != nil {
		logger.Log("Search error", err.Error())
		return nil, err
	}
	return rep.(*searchMovies.SearchResponse), nil
}

func (g *grpcServer) HealthCheck(ctx context.Context, r *searchMovies.HealthCheckRequest) (*searchMovies.HealthCheckResponse, error) {
	_, rep, err := g.heathCheck.ServeGRPC(ctx, r)
	if err != nil {
		logger.Log("HealthCheck error", err.Error())
		return nil, err
	}
	return rep.(*searchMovies.HealthCheckResponse), nil
}

func decodeGRPCSearchRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req, ok := grpcReq.(*searchMovies.SearchRequest)
	if !ok {
		logger.Log("decodeGRPCSearchRequest", "decode request error")
		return nil, errors.New("searchMovies decode request error")
	}

	return model.SearchRequest{Pagination: req.Pagination, SearchWord: req.SearchWord}, nil
}

func decodeGRPCHealthCheckRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	return model.HealthCheckRequest{}, nil
}

func decodeGRPCSearchResponse(_ context.Context, grpcResp interface{}) (interface{}, error) {
	resp, ok := grpcResp.(model.SearchResponse)
	if !ok {
		logger.Log("decodeGRPCSearchResponse", "decode response error")
		return nil, errors.New("searchMovies decode response error")
	}

	var movies []*searchMovies.Movie
	for _, val := range resp.Movies {
		movie := searchMovies.Movie{
			Title:  val.Title,
			Year:   val.Year,
			ImdbID: val.ImdbID,
			Type:   val.Type,
			Poster: val.Poster,
		}
		movies = append(movies, &movie)
	}

	return &searchMovies.SearchResponse{Movies: movies, Err: resp.Err}, nil
}

func decodeGRPCHealthCheckResponse(_ context.Context, grpcResp interface{}) (interface{}, error) {
	resp, ok := grpcResp.(model.HealthCheckResponse)
	if !ok {
		logger.Log("decodeGRPCHealthCheckResponse", "decode response error")
		return nil, errors.New("healthcheck decode response error")
	}
	return &searchMovies.HealthCheckResponse{Code: resp.Code, Err: resp.Err}, nil
}
