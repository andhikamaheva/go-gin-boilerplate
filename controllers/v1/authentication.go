package authentication

import (
	"github.com/andhikamaheva/TheArisanBot/helpers"
	"github.com/gin-gonic/gin"
)

// Info function
func Info(c *gin.Context) {
	meta.Response(c, "success", 200, meta.Message{"a", "a", "a"})
}
