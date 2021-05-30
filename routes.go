package main

func initializeRouter() {
	router.GET("/", showIndexPage)
	//router.GET("/articles")
	router.GET("/article/:id", showArticle)
	router.POST("/article", createArticle)
}
