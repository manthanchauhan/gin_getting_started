package controllers

import (
	"github.com/manthanchauhan/gin_getting_started/models"
	"gorm.io/gorm"
)

type ArticleGetController struct {
	Filters map[string]interface{}
	DB      *gorm.DB
}

func (contr ArticleGetController) getArticle() models.Article {
	var article models.Article
	contr.DB.Where(contr.Filters).Last(&article)

	return article
}
