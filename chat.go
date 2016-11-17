package oi

import (
	"github.com/jinzhu/gorm"
)

type Chat struct {
	ID       int
	Messages []Message
	gorm.Model
}

type Message struct {
	ID      int
	Message string
	Subject string
	ChatID  int
	gorm.Model
}

func Connect(name string) *gorm.DB {
	db, err := gorm.Open("postgres", name)
	if err != nil {
		return err
	}
	return db
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&Chat{})
	db.AutoMigrate(&Message{})
}
