package model

import (
	"math/rand"

	"gorm.io/gorm"
)

type Link struct {
	gorm.Model
	UserID uint `gorm:"not null"`
	//TODO: deal with CASCADE deletion
	Stats []Stat `gorm:"constraint:OnDelete:CASCADE;"`
	Url   string `json:"url" gorm:"not null"`
	Hash  string `json:"hash" gorm:"uniqueIndex,not null"`
}

func NewLink(url string, userId uint) *Link {
	return &Link{
		UserID: userId,
		Url:    url,
		Hash:   randStringRunes(10),
	}
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
