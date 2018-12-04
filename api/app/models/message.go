package models

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

type Message struct {
	ID        int       `json:"id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func Messages(db *gorm.DB) ([]Message, error) {
	var messages []Message
	if db == nil {
		return nil, fmt.Errorf("db is nil")
	}
	if err := db.Find(&messages).Error; err != nil {
		return nil, err
	}

	return messages, nil
}
