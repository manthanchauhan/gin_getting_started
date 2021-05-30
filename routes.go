package main

func initializeRouter() {
	router.GET("/", showIndexPage)
	router.GET("/article/:id", showArticle)
}
