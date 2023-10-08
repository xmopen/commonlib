package ginhelper

import "github.com/gin-gonic/gin"

// RegisterRequest return gin.handler
func RegisterRequest() func(c *gin.Context) {
	return func(c *gin.Context) {
		xlog := Log(c)
		xlog.Infof("request:[%+v]", c.Request.URL.Host+c.Request.URL.Path+c.Request.URL.RawQuery)
		c.Next()
	}
}
