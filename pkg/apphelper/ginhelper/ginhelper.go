package ginhelper

import (
	"github.com/gin-gonic/gin"
	"github.com/xmopen/golib/pkg/utils"
	"github.com/xmopen/golib/pkg/xlogging"
)

// Log return an xlogging instance for gin.tcontext
func Log(c *gin.Context) *xlogging.Entry {
	xlogItr, ok := c.Get("xlog")
	if ok {
		xlog, ok := xlogItr.(*xlogging.Entry)
		if ok {
			return xlog
		}
	}
	xlog := xlogging.Tag(utils.UUID())
	c.Set("xlog", xlog)
	return xlog
}
