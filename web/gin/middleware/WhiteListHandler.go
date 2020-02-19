package tymiddleware

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	tywhitelist "gocommon/web/whitelist"
	"net/http"
)

// IPWhiteListHandler 白名单验证
func IPWhiteListHandler(whiteList []string) gin.HandlerFunc {
	wlst := tywhitelist.New(whiteList...)
	return func(c *gin.Context) {
		clientIP := c.ClientIP()
		if err := wlst.Allow(clientIP); err != nil {
			log.WithFields(log.Fields{"clientIP": clientIP}).Error(err)
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": wlst.ErrMsg()})
			return
		}

		c.Next()
	}
}
