package main

import "LearnGoGin/Controller"

func initializeRoutes() {
	router.GET("/", Controller.ShowIndexPage)
	router.GET("/article/view/:article_id", Controller.GetArticle)
}
