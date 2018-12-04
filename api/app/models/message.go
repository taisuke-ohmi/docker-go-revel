package main

import "github.com/jinzhu/gorm"

type Message struct {
	ID      int
	Content string
}

func Messages(db *gorm.DB) ([]Message, error) {
	var messages []Message
	if err := db.Find(&messages).Error; err != nil {
		return nil, err
	}

	return messages, nil
}
