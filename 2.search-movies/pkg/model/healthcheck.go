package model

type HealthCheckRequest struct{}

type HealthCheckResponse struct {
	Code int64  `json:"status"`
	Err  string `json:"err,omitempty"`
}
