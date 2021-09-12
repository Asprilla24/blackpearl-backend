package model

type HealthCheckRequest struct{}

type HealthCheckResponse struct {
	Code int64  `json:"code"`
	Err  string `json:"err,omitempty"`
}
