package db_managers

import (
	"github.com/manthanchauhan/gin_getting_started/models"
	"gorm.io/gorm"
)

type ArticleDBManager struct {
	DB *gorm.DB
}

func (dbMgr ArticleDBManager) CreateArticle(article models.Article) {
	dbMgr.DB.Create(&article)
}
