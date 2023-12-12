package models

import "time"

type Node struct {
	Id               uint   `gorm:"primary_key" json:"id"`
	IP               string `json:"ip"`
	Port             string `json:"port"`
	Score            int64  `gorm:"default:100" json:"score"`
	Priority         int64  `gorm:"default:0" json:"priority"`
	CreatedAtUnix    int64  `gorm:"autoCreateTime" json:"created_at"`
	LastResponseUnix int64  `gorm:"autoCreateTime" json:"last_response"`
	LastRequestUnix  int64  `gorm:"autoCreateTime" json:"last_request"`
	IsActive         bool   `gorm:"default:true" json:"is_active"`

	RequestCount              uint      `gorm:"-"`
	NextRequestCountResetTime time.Time `gorm:"-"`
	PingArray                 []int64   `gorm:"-" json:"ping_array"`
	ResponseTimeArray         []int64   `gorm:"-" json:"response_time_array"`
}
