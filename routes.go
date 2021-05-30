package main

func initializeRouter() {
	router.GET("/articles", listAllArticles)
	//router.GET("/article/:id", showArticle)
	router.POST("/article", createArticle)
}
