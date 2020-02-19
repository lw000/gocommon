package tymiddleware

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"gocommon/web/regionFiltering"
	"net/http"
)

// IP区域过滤
func RegionFilteringHandler(ip2regiondb string, regions ...string) gin.HandlerFunc {
	regionfilter := regionFiltering.New()
	if err := regionfilter.Load(ip2regiondb); err != nil {
		log.Panic(err)
	}

	return func(c *gin.Context) {
		clientIP := c.ClientIP()
		err := regionfilter.Allow(clientIP, regions...)
		if err != nil {
			log.WithFields(log.Fields{"clientIP": clientIP}).Error(err)
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"c": 0, "m": err.Error()})
			return
		}

		c.Next()
	}
}
