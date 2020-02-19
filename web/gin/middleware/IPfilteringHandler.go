package tymiddleware

import (
	"github.com/gin-gonic/gin"
	"github.com/lw000/gocommon/web/ipfiltering"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func IPfilteringHandler(ip2regiondb string, AbroadAccess bool, whiteList ...string) gin.HandlerFunc {
	ipfilter := ipfiltering.New(AbroadAccess, whiteList...)
	if err := ipfilter.Load(ip2regiondb); err != nil {
		log.Panic(err)
	}

	return func(c *gin.Context) {
		clientIP := c.ClientIP()
		allow, err := ipfilter.Allow(clientIP)
		if err != nil {
			log.WithFields(log.Fields{"clientIP": clientIP}).Error(err)
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"c": 0, "m": err.Error()})
			return
		}

		if !allow {
			log.WithFields(log.Fields{"clientIP": clientIP}).Error(err)
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"c": 0, "m": err.Error()})
			return
		}

		c.Next()
	}
}
