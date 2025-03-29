package link

import (
	"math/rand"

	"gorm.io/gorm"
)

type Link struct {
	gorm.Model
	Url  string `json:"url" gorm:"not null"`
	Hash string `json:"hash" gorm:"uniqueIndex,not null"`
}

func NewLink(url string) *Link {
	return &Link{
		Url:  url,
		Hash: randStringRunes(10),
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
