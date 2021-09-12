package transports

import (
	"context"

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
		return nil, err
	}
	return rep.(*searchMovies.SearchResponse), nil
}

func (g *grpcServer) HealthCheck(ctx context.Context, r *searchMovies.HealthCheckRequest) (*searchMovies.HealthCheckResponse, error) {
	_, rep, err := g.heathCheck.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return rep.(*searchMovies.HealthCheckResponse), nil
}

func decodeGRPCSearchRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*searchMovies.SearchRequest)

	return model.SearchRequest{Pagination: req.Pagination, SearchWord: req.SearchWord}, nil
}

func decodeGRPCHealthCheckRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	return model.HealthCheckRequest{}, nil
}

func decodeGRPCSearchResponse(_ context.Context, grpcResp interface{}) (interface{}, error) {
	resp := grpcResp.(model.SearchResponse)
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
	resp := grpcResp.(model.HealthCheckResponse)
	return &searchMovies.HealthCheckResponse{Code: resp.Code, Err: resp.Err}, nil
}
