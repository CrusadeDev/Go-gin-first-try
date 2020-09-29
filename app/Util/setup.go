package Util

import "github.com/gin-gonic/gin"

var router *gin.Engine

func InitRouter() *gin.Engine {
	router = gin.Default()

	return router
}

func GetRouter() *gin.Engine {
	return router
}
