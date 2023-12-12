package options

import (
	"time"

	"github.com/VDiPaola/api-network-server/helpers"
)

var options = helpers.OptionsType{
	RateLimitAmount:     300,
	RateLimitInterval:   time.Duration(time.Minute * 5),
	MaxResponseDuration: time.Duration(time.Millisecond * 100),
	NodeRedundencyCount: 2,
	HealthCheckDuration: time.Duration(time.Minute * 15),
	CacheUpdateInterval: time.Duration(time.Minute * 5),
	HasCache:            true,
	ScoreThreshold:      10,
}

func Set(newOptions helpers.OptionsType) {
	options = newOptions
}

func Get() helpers.OptionsType {
	return options
}
