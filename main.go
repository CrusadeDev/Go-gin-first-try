package main

import (
	"LearnGoGin/app/Route"
	"LearnGoGin/app/Util"
)

func main() {
	router := Util.InitRouter()

	Route.InitRoutes()

	Util.ConnectToDatabase()

	router.Run()
}
