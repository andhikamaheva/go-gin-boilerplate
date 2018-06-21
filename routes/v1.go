package v1

import (
	"github.com/andhikamaheva/TheArisanBot/controllers/v1"
	"github.com/gin-gonic/gin"
)

// Routes for API Version 1 (v1)
func Routes(route *gin.Engine) {
	v1 := route.Group("/v1")
	{
		v1.GET("/authentication/info", authentication.Info)
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(400, gin.H{
				"message": "pong",
			})
		})
	}
}
