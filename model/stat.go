package model

import (
	"gorm.io/gorm"
)

type Stat struct {
	gorm.Model
	LinkID uint `gorm:"uniqueIndex:idx_stats_link_id,not null"`
	Click  uint `json:"click" gorm:"not null"`
}

func NewStat(linkId uint) *Stat {
	return &Stat{
		LinkID: linkId,
		Click:  1,
	}
}
