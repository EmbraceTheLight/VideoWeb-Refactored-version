package gsserver

import "github.com/gin-gonic/gin"

type VideoDownloadFileServer interface {
	GetMpd(c *gin.Context)
	GetSegments(c *gin.Context)
	GetVideoCover(c *gin.Context)
}
