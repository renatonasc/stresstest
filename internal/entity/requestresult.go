package entity

import "time"

// RequestResult is the result of a request
type RequestResult struct {
	StatusCode int
	Body       string
	Err        error
	Request    Request
	Duration   time.Duration
	TimeStart  time.Time
	TimeEnd    time.Time
}
