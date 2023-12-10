package models

import "time"

type Node struct {
	Id                  uint   `gorm:"primary_key" json:"id"`
	IP                  string `json:"ip"`
	Score               int64  `gorm:"default:100" json:"score"`
	Priority            int64  `gorm:"default:0" json:"priority"`
	CreatedAtUnix       int64  `gorm:"autoCreateTime" json:"created_at"`
	LastResponseUnix    int64  `gorm:"autoCreateTime" json:"last_response"`
	LastRequestUnix     int64  `gorm:"autoCreateTime" json:"last_request"`
	AveragePing         int64  `json:"average_ping"`
	AverageResponseTime int64  `json:"average_response_time"`
	IsActive            bool   `gorm:"default:true" json:"is_active"`

	RequestCount              uint      `gorm:"-"`
	NextRequestCountResetTime time.Time `gorm:"-"`
}
