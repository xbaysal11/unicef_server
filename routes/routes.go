package routes

import (
	"github.com/gin-gonic/gin"
)

func AddRoutes(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ping",
		})
	})
}
