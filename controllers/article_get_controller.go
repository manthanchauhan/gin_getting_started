package controllers

import (
	"github.com/manthanchauhan/gin_getting_started/models"
	"gorm.io/gorm"
)

type ArticleGetController struct {
	filters map[string]interface{}
	db      *gorm.DB
}

func (contr ArticleGetController) getArticle() models.Article {
	var article models.Article
	contr.db.Where(contr.filters).Last(&article)

	return article
}
