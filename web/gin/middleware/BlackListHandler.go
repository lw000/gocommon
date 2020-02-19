package tymiddleware

import (
	"github.com/gin-gonic/gin"
	"github.com/lw000/gocommon/web/blacklist"
	log "github.com/sirupsen/logrus"
	"net/http"
)

// IPBlackListHandler 黑名单单验证
func IPBlackListHandler(blackList []string) gin.HandlerFunc {
	blst := tyblacklist.New(blackList...)
	return func(c *gin.Context) {
		clientIP := c.ClientIP()
		if err := blst.Deny(clientIP); err != nil {
			log.WithFields(log.Fields{"clientIP": clientIP}).Error(err)
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": blst.ErrMsg()})
			return
		}

		c.Next()
	}
}
