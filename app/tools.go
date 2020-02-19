package tyapp

import (
	"github.com/gin-gonic/gin"
)

func ReplyJSON(c *gin.Context, code int, d interface{}) {
	c.JSON(code, d)
}

func ReplyYAML(c *gin.Context, code int, d interface{}) {
	c.YAML(code, d)
}
func ReplyXML(c *gin.Context, code int, d interface{}) {
	c.XML(code, d)
}
