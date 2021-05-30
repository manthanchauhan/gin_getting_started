package main

import (
	"github.com/gin-gonic/gin"
	"github.com/manthanchauhan/gin_getting_started/db"
	"github.com/manthanchauhan/gin_getting_started/middlewares"
)

var router *gin.Engine

func main() {
	database := db.Connect()

	router = gin.Default()
	router.Use(middlewares.SetDBToContext(database))

	initializeRouter()

	router.Run()
}
