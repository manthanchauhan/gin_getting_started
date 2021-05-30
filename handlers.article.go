package main

import (
	"github.com/gin-gonic/gin"
	"github.com/manthanchauhan/gin_getting_started/controllers"
	dbModule "github.com/manthanchauhan/gin_getting_started/db"
	"github.com/manthanchauhan/gin_getting_started/models"
)

//func showIndexPage(c *gin.Context) {
//	articles := getAllArticles()
//	c.JSON(http.StatusOK, articles)
//}
//
//func showArticle(c *gin.Context) {
//	articleId, err := strconv.Atoi(c.Param("id"))
//
//	if err != nil {
//		c.AbortWithStatus(http.StatusBadRequest)
//	}
//
//	article_, err := getArticleById(articleId)
//
//	if err != nil {
//		c.AbortWithStatus(http.StatusNotFound)
//	}
//
//	c.JSON(http.StatusOK, *article_)
//}

func createArticle(c *gin.Context) {
	article_ := models.Article{}
	err := c.Bind(&article_)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	db := dbModule.DBInstance(c)
	db.Create(&article_)

	c.JSON(201, article_)
}

func listAllArticles(c *gin.Context) {
	db := dbModule.DBInstance(c)

	articles := controllers.ListArticles(db)

	c.JSON(200, articles)
}
