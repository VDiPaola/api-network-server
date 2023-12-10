package helpers

import (
	"net/http"
	"time"
)

type RequestCallbackType func(response *http.Response, err error)
type ResponseCallbackType func(response *http.Response, err error)

type OptionsType struct {
	RateLimitAmount   uint //amount of request per RateLimitInterval
	RateLimitInterval time.Duration
	//MaxRequestsPerMinute uint
	MaxResponseDuration time.Duration //max time a node is allowed to respond in
	NodeRedundencyCount uint          //how many nodes to ask a request to simultaneously
	HealthCheckDuration time.Duration //how long to wait before considering a node inactive
	CacheUpdateInterval time.Duration
	HasCache            bool
	ScoreThreshold      int64
}

// var options = Options{
// 	RateLimitAmount:     300,
// 	RateLimitInterval:   time.Duration(time.Minute * 5),
// 	MaxResponseDuration: time.Duration(time.Millisecond * 100),
// 	NodeRedundencyCount: 2,
// 	HealthCheckDuration: time.Duration(time.Minute * 15),
// 	CacheUpdateInterval: time.Duration(time.Minute * 5),
// 	HasCache:            true,
// 	ScoreThreshold:      10,
// }
