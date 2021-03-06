package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/imatheus-lucas/golang_api/controllers"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		books := main.Group("books")
		{
			books.GET("/:id", controllers.ShowBook)
			books.POST("/", controllers.CreateBook)
			books.GET("/", controllers.ShowAllBooks)
		}
	}

	return router
}
