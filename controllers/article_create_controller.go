package controllers

import (
	"github.com/manthanchauhan/gin_getting_started/db_managers"
	"github.com/manthanchauhan/gin_getting_started/models"
	"gorm.io/gorm"
)

type ArticleCreateContr struct {
	DB         *gorm.DB
	NewArticle *models.Article
}

func (contr ArticleCreateContr) CreateArticle() (bool, error) {
	dbMgr := db_managers.ArticleDBManager{DB: contr.DB}
	success, err := dbMgr.CreateArticle(contr.NewArticle)

	return success, err
}
