package model

import "errors"

type Article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var articleList = []Article{
	{ID: 1, Title: "Article 1", Content: "Article 1 body"},
	{ID: 2, Title: "Article 2", Content: "Article 2 body"},
}

func GetAllArticles() []Article {
	return articleList
}

func GetArticleById(articleId int) (*Article, error) {
	for _, a := range articleList {
		if a.ID == articleId {
			return &a, nil
		}
	}

	return nil, errors.New("article not found")
}
