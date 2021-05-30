package controllers

import (
	"github.com/manthanchauhan/gin_getting_started/models"
	"gorm.io/gorm"
)

func ListArticles(db *gorm.DB) {
	var articles []models.Article

	result := db.Find(&articles)
	println(result)
}
