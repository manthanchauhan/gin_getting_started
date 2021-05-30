package db_managers

import (
	"github.com/manthanchauhan/gin_getting_started/models"
	"gorm.io/gorm"
)

type ArticleDBManager struct {
	DB *gorm.DB
}

func (dbMgr ArticleDBManager) CreateArticle(article *models.Article) (bool, error) {
	result := dbMgr.DB.Create(article)

	if result.Error != nil {
		return false, result.Error
	}

	return true, nil
}
