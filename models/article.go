package models

import (
	"time"
)

type Article struct {
	// gorm.Model
	ID           uint `gorm:"primary_key"`
	Title        string
	Description  string
	Content      string
	ArticleUrl   string
	ImageUrl     string
	ResourceName string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
