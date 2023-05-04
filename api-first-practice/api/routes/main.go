package routes

import (
	"api/api/controllers"

	"github.com/gin-gonic/gin"
)

func AppRoutes(router *gin.Engine) *gin.RouterGroup {
	tweetController := controllers.NewTweetController()

	route := router.Group("/v1")
	{
		route.GET("/tweets", tweetController.FindAll)
		route.POST("/tweet", tweetController.Create)
		route.DELETE("/tweet/:id", tweetController.Delete)
	}

	return route
}
