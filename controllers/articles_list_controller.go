package controllers

import (
	"github.com/manthanchauhan/gin_getting_started/models"
	"gorm.io/gorm"
)

func listArticles(db *gorm.DB) {
	var articles []models.Article

	result := db.Find(&articles)
	println(result)
}
