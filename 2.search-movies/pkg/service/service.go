package service

import (
	"context"
	"search-movies/pkg/model"
)

type Service interface {
	Search(ctx context.Context, pagination int64, searchWord string) ([]model.Movie, error)
	HealthCheck(ctx context.Context) (int64, error)
}
