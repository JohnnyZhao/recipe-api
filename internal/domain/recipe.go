package domain

import (
	"time"
)

type Recipe struct {
	ID          uint      `json:"id"`
	Title       string    `gorm:"type:varchar(100);not null" json:"title"`
	MakingTime  string    `gorm:"type:varchar(100);not null" json:"making_time"`
	Serves      string    `gorm:"type:varchar(100);not null" json:"serves"`
	Ingredients string    `gorm:"type:varchar(300);not null" json:"ingredients"`
	Cost        int       `gorm:"not null" json:"cost"`
	CreatedAt   time.Time `gorm:"not null" json:"-"`
	UpdatedAt   time.Time `gorm:"not null" json:"-"`
}
