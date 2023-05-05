package models

import (
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	"log"
)

type Article struct {
	gorm.Model
	ID uint
	Title string
	Description string
	Content string
	ArticleUrl string
	ImageUrl string
	ResourceName string
}

func NewArticle(title string, description string, content string, articleUrl string, imageUrl string, resourceName string) *Article {
	return &Article {
		Title: title,
		Description: description,
		Content: content,
		ArticleUrl: articleUrl,
		ImageUrl: imageUrl,
		ResourceName: resourceName,
	}
}

func (a *Article) Create() {
	db, err := gorm.Open(mysql.Open("news_article"), &gorm.Config{})
	if err != nil {
		log.Println(err)
	}

	db.Create(&a)
}
