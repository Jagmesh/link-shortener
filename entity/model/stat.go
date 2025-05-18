package model

import (
	"gorm.io/gorm"
	"time"
)

type Stat struct {
	gorm.Model
	LinkID  uint   `gorm:"uniqueIndex:idx_stats_link_day;not null"`
	Click   uint   `json:"click" gorm:"not null"`
	DayDate string `gorm:"uniqueIndex:idx_stats_link_day;not null;type:DATE"`
}

func NewStat(linkId uint) *Stat {
	return &Stat{
		LinkID:  linkId,
		Click:   1,
		DayDate: time.Now().Format(time.DateOnly),
	}
}
