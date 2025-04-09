package gs

import (
	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/log"
	"path/filepath"
	"util/helper"
	"util/helper/file"
	"vw_gateway"
	"vw_gateway/internal/biz/videobiz"
	"vw_gateway/internal/pkg/ecode/errdef/uerr"
)

const dashPath = "dash"

type VideoDownloadFileService struct {
	vinfoUsecase *videobiz.VideoInfoUsecase
	logger       *log.Helper
}

func NewVideoDownloadFileService(usecase *videobiz.VideoInfoUsecase, logger log.Logger) *VideoDownloadFileService {
	return &VideoDownloadFileService{
		vinfoUsecase: usecase,
		logger:       log.NewHelper(logger),
	}
}

// GetMpd
// Params:
// - path_params: video_id, author_id. They are used to compute the .mpd path of the video.
func (s *VideoDownloadFileService) GetMpd(c *gin.Context) {
	videoId := c.Param("video_id")
	authorId := c.Param("author_id")
	mpdPath := filepath.Join(vw_gateway.ResourcePath, authorId, videoId, dashPath, "output.mpd")
	isExists := file.CheckIfFileExist(mpdPath)
	if !isExists {
		c.JSON(200, gin.H{
			"msg": "请求的.mpd文件路径不存在!",
		})
	}
	c.File(mpdPath)
}

// GetSegments
// Params:
// - query_params: file_path. // TODO: The param may be changed later.
func (s *VideoDownloadFileService) GetSegments(c *gin.Context) {
	filePath := c.Query("file_path")
	isExists := file.CheckIfFileExist(filePath)
	if !isExists {
		c.JSON(200, gin.H{
			"message": helper.HandleError(uerr.ErrSegmentFileNotFound),
		})
	}
	c.File(filePath)
}

// GetVideoCover
// Params:
// - path_params: video_id, author_id. They are used to compute the cover path of the video.
func (s *VideoDownloadFileService) GetVideoCover(c *gin.Context) {
	videoId := c.Param("video_id")
	authorId := c.Param("author_id")
	targetDir := filepath.Join(vw_gateway.ResourcePath, authorId, videoId)
	coverPath, err := file.NewFileSearcher().Find(targetDir, "cover")
	if err != nil {
		c.JSON(200, gin.H{
			"message": helper.HandleError(uerr.ErrCoverFileNotFound, err),
		})
	}
	c.File(coverPath)
}
