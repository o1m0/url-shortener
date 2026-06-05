package models

import (
	"time"

	"gorm.io/gorm"
)

type URL struct {
	gorm.Model  `json:"-"`
	ID          uint       `gorm:"primarykey" json:"id"`
	CreatedAt   time.Time  `json:"created_at"`
	OriginalURL string     `gorm:"not null" json:"original_url"`
	ShortCode   string     `gorm:"uniqueIndex;not null" json:"short_code"`
	UserID      *uint      `json:"user_id"`
	AccessCount int        `gorm:"default:0" json:"access_count"`
	ExpiresAt   *time.Time `json:"expires_at"`
}

type User struct {
	gorm.Model
	Email    string `gorm:"uniqueIndex;not null" json:"email"`
	Password string `gorm:"not null" json:"-"`
	URLs     []URL  `gorm:"foreignKey:UserID" json:"urls"`
}
