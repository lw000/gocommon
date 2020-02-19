package tymiddleware

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/unrolled/secure"
)

// TlsHandler tls开启
func TlsHandler(host string) gin.HandlerFunc {
	return func(c *gin.Context) {
		secureMiddleware := secure.New(secure.Options{
			SSLRedirect: true,
			SSLHost:     host,
		})

		er := secureMiddleware.Process(c.Writer, c.Request)
		if er != nil {
			log.Error(er)
			c.Abort()
			return
		}

		c.Next()
	}
}
