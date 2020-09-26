package Controller

import (
	"LearnGoGin/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ShowIndexPage(c *gin.Context) {
	particles := model.GetAllArticles()
	c.HTML(http.StatusOK, "index.html", gin.H{"title": "Home Page", "payload": particles})
}
