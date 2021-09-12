package transports

import (
	"context"
	"encoding/json"
	"net/http"
	"os"

	"search-movies/pkg/endpoints"
	"search-movies/pkg/model"
	"search-movies/pkg/util"

	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
)

func NewHTTPHandler(ep endpoints.Set) http.Handler {
	m := http.NewServeMux()

	m.Handle("/healthz", httptransport.NewServer(
		ep.HealthCheckEndpoint,
		decodeHTTPHealthCheckRequest,
		encodeResponse,
	))
	m.Handle("/search", httptransport.NewServer(
		ep.SearchEndpoint,
		decodeHTTPSearchRequest,
		encodeResponse,
	))

	return m
}

func decodeHTTPSearchRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req model.SearchRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		logger.Log("decodeHTTPSearchRequest", err.Error())
		return nil, err
	}
	return req, nil
}

func decodeHTTPHealthCheckRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req model.HealthCheckRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		logger.Log("decodeHTTPHealthCheckRequest", err.Error())
		return nil, err
	}
	return req, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(error); ok && e != nil {
		encodeError(ctx, e, w)
		return nil
	}
	return json.NewEncoder(w).Encode(response)
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	switch err {
	case util.ErrBadRequest:
		w.WriteHeader(http.StatusBadRequest)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

var logger log.Logger

func init() {
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "transports", log.DefaultTimestampUTC)
}
