package models

import "time"

type Article struct {
	ID        uint      `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
