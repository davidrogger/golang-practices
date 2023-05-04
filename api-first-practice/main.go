package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	app.GET("/v1", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "online",
		})
	})

	app.Run("localhost:3002")
}
