package main

import (
	"github.com/gin-gonic/gin"
	"github.com/manthanchauhan/gin_getting_started/db"
)

var router *gin.Engine

func main() {
	database := db.Connect()
	router = gin.Default()

	initializeRouter()

	router.Run()
}
