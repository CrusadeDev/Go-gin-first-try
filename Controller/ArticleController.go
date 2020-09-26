package Controller

import (
	"LearnGoGin/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetArticle(c *gin.Context) {
	articleId, err := strconv.Atoi(c.Param("article_id"))

	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
	}

	article, err := model.GetArticleById(articleId)

	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
	}

	c.HTML(
		// Set the HTTP status to 200 (OK)
		http.StatusOK,
		// Use the article.html template
		"article.html",
		// Pass the data that the page uses
		gin.H{
			"title":   article.Title,
			"payload": article,
		},
	)
}
