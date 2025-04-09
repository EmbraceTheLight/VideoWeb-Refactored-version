package domain

import (
	videoinfov1 "vw_video/api/v1/videoinfo"
)

type Records struct {
	CntBarrages  uint32
	CntShares    uint32
	CntViewed    uint32
	CntFavorites uint32
}
type VideoDetail struct {
	VideoId       int64
	PublisherId   int64
	PublisherName string
	Title         string
	Description   string
	VideoPath     string
	Classes       []string
	Tags          []string
	Hot           int64
	Records       *Records
	Duration      string
	CoverPath     string
}

// NewVideoDetail creates a new VideoDetail object from video grpc service GetVideoInfo's response.
func NewVideoDetail(videoInfo *videoinfov1.GetVideoInfoResp) *VideoDetail {
	return &VideoDetail{
		VideoId:       videoInfo.VideoDetail.VideoId,
		PublisherId:   videoInfo.VideoDetail.PublisherId,
		PublisherName: videoInfo.VideoDetail.PublisherName,
		Title:         videoInfo.VideoDetail.Title,
		Description:   videoInfo.VideoDetail.Description,
		VideoPath:     videoInfo.VideoDetail.VideoPath,
		Classes:       videoInfo.VideoDetail.Classes,
		Tags:          videoInfo.VideoDetail.Tags,
		Hot:           videoInfo.VideoDetail.Hot,
		Records: &Records{
			CntBarrages:  videoInfo.VideoDetail.Records.CntBarrages,
			CntShares:    videoInfo.VideoDetail.Records.CntShares,
			CntViewed:    videoInfo.VideoDetail.Records.CntViewed,
			CntFavorites: videoInfo.VideoDetail.Records.CntFavorites,
		},
		Duration:  videoInfo.VideoDetail.Duration,
		CoverPath: videoInfo.VideoDetail.CoverPath,
	}
}

// NewVideoDetails creates a new VideoDetail object from video grpc service GetVideoList's response.
func NewVideoDetails(videoInfo *videoinfov1.GetVideoListResp) []*VideoDetail {
	ret := make([]*VideoDetail, len(videoInfo.VideoDetails))

	for i, videoDetail := range videoInfo.VideoDetails {
		ret[i] = &VideoDetail{
			VideoId:       videoDetail.VideoId,
			PublisherId:   videoDetail.PublisherId,
			PublisherName: videoDetail.PublisherName,
			Title:         videoDetail.Title,
			Description:   videoDetail.Description,
			VideoPath:     videoDetail.VideoPath,
			Classes:       videoDetail.Classes,
			Tags:          videoDetail.Tags,
			Hot:           videoDetail.Hot,
			Records: &Records{
				CntBarrages:  videoDetail.Records.CntBarrages,
				CntShares:    videoDetail.Records.CntShares,
				CntViewed:    videoDetail.Records.CntViewed,
				CntFavorites: videoDetail.Records.CntFavorites,
			},
			Duration:  videoDetail.Duration,
			CoverPath: videoDetail.CoverPath,
		}
	}
	return ret
}
