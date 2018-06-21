package meta

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// Message : Response message
type Message struct {
	Param   string
	Message string
	Value   string
}

// Response of Meta
func Response(c *gin.Context, responseType string, code int, messages []Message) {
	fmt.Println(messages)
	c.JSON(code, gin.H{
		"type": responseType,
		"code": code,
		//"messages": messages,
	})
}
