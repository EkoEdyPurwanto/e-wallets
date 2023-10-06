package req

import "time"

type LoggingRequest struct {
	StartTime  time.Time
	EndTime    time.Duration
	StatusCode int
	ClientIP   string
	Method     string
	Path       string
	UserAgent  string
}
