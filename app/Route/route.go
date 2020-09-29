package Route

import (
	"LearnGoGin/app/Handler"
	"LearnGoGin/app/Util"
)

func InitRoutes() {
	router := Util.GetRouter()

	router.GET("/user/:id", Handler.FindByIdHandler)
}
