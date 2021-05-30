package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/manthanchauhan/gin_getting_started/controllers"
	dbModule "github.com/manthanchauhan/gin_getting_started/db"
	"github.com/manthanchauhan/gin_getting_started/models"
	"strconv"
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

	contr := controllers.ArticleCreateContr{NewArticle: &article_, DB: db}
	_, err = contr.CreateArticle()

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, article_)
}

func listAllArticles(c *gin.Context) {
	db := dbModule.DBInstance(c)

	articles := controllers.ListArticles(db)

	c.JSON(200, articles)
}

func getArticleById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.AbortWithStatus(400)
	}

	if id == 0 {
		c.AbortWithError(400, errors.New("invalid article id"))
		return
	}

	db := dbModule.DBInstance(c)

	filters := map[string]interface{}{"ID": id}

	contr := controllers.ArticleGetController{Filters: filters, DB: db}
	article := contr.GetArticle()

	c.JSON(200, article)
}
