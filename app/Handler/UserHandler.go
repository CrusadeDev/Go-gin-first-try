package Handler

import (
	"LearnGoGin/app/Model"
	"github.com/gin-gonic/gin"
)

type findByIdRequest struct {
	Id int `uri:"id" binding:required"`
}

func FindByIdHandler(c *gin.Context) {
	request := findByIdRequest{}

	err := c.BindUri(&request)

	if err != nil {
		err := c.AbortWithError(400, err)
		c.JSON(400, err)

		return
	}

	user, err := Model.FindById(request.Id)

	if err != nil {
		err := c.AbortWithError(404, err)
		c.JSON(400, err)

		return
	}

	c.JSON(200, user)
}
