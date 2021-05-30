package controllers

import (
	"github.com/manthanchauhan/gin_getting_started/models"
	"gorm.io/gorm"
)

func ListArticles(db *gorm.DB) []models.Article {
	var articles []models.Article

	db.Find(&articles)
	return articles
}
