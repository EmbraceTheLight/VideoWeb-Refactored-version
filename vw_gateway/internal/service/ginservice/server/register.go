package ginserver

import (
	"github.com/gin-gonic/gin"
	"vw_gateway/internal/service/ginservice/server/servers"
)

func RegisterVideoFileDownloadHTTPServer(ge *gin.Engine, srv gsserver.VideoDownloadFileServer) *gin.Engine {
	ge.GET("/:video_id/mpd", srv.GetMpd)
	ge.GET("/:video_id/segments", srv.GetSegments)
	ge.GET("/:video_id/cover", srv.GetVideoCover)
	return ge
}
