package tymiddleware

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"levi/gocommon/web/hostBinding"
	"net/http"
)

// HostBindingHandler 域名绑定认证
func HostBindingHandler(hosts []string) gin.HandlerFunc {
	binding := hostBinding.New()
	binding.Binding(hosts...)
	return func(c *gin.Context) {
		if err := binding.Allow(c.Request.Host); err != nil {
			log.WithFields(log.Fields{"clientIP": c.ClientIP(), "host": c.Request.Host}).Error(err)
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.Next()
	}
}
