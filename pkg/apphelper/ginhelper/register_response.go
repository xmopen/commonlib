package ginhelper

import "github.com/gin-gonic/gin"

// RegisterResponse register response middleware for gin.tcontext
func RegisterResponse() func(c *gin.Context) {
	return func(c *gin.Context) {
	}
}
