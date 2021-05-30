package main

import "errors"

type article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var articleList = []article{
	{1, "Article 1", "Article 1 Content"},
	{2, "Article 2", "Article 2 Content"},
}

func getAllArticles() []article {
	return articleList
}

func getArticleById(id int) (*article, error) {
	for _, article := range articleList {
		if article.ID == id {
			return &article, nil
		}
	}

	return nil, errors.New("article not found")
}
