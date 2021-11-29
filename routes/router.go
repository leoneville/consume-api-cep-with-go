package routes

import (
	"cep/controllers"

	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		cep := main.Group("cep")
		{
			cep.GET("/:cep", controllers.Cep)
		}
	}
	return router
}
