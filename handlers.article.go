package main

import (
	"github.com/gin-gonic/gin"
	dbModule "github.com/manthanchauhan/gin_getting_started/db"
	"github.com/manthanchauhan/gin_getting_started/models"
	"net/http"
	"strconv"
)

func showIndexPage(c *gin.Context) {
	articles := getAllArticles()
	c.JSON(http.StatusOK, articles)
}

func showArticle(c *gin.Context) {
	articleId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}

	article_, err := getArticleById(articleId)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}

	c.JSON(http.StatusOK, *article_)
}

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
