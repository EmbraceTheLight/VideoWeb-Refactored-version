package videodata

import (
	"context"
	stderr "errors"
	"github.com/dtm-labs/dtm/client/dtmgrpc"
	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/proto"
	"util/getid"
	"util/resolver"
	interactv1 "vw_gateway/api/v1/video/video_interact"
	"vw_gateway/internal/biz/videobiz"
	userinfov1 "vw_user/api/v1/userinfo"
	videointeractv1 "vw_video/api/v1/interact"
	videoinfov1 "vw_video/api/v1/videoinfo"
)

type interactRepo struct {
	data   *Data
	logger *log.Helper
}

func NewInteractRepo(data *Data, logger log.Logger) videobiz.InteractRepo {
	return &interactRepo{
		data:   data,
		logger: log.NewHelper(logger),
	}
}

type sagaMember struct {
	method           string
	compensateMethod string
	payload          proto.Message
}

func (i *interactRepo) FavoriteVideo(ctx context.Context, userId, favoriteId, videoId int64) (bool, error) {
	// 1. Get the current status of the user-video
	origin, err := i.data.videoInteractClient.GetUserVideoStatus(ctx, &videointeractv1.GetUserVideoStatusReq{
		VideoId: videoId,
		UserId:  userId,
	})
	if err != nil {
		return false, err
	}

	gid, saga := newSaga(ctx, i.data.dtmServerAddr)
	i.logger.Infof("[FavoriteVideo] gid: %s", gid)

	// 2. Get the address of the video service. format --> ip:port
	addr, err := resolver.GetRandomAddr(ctx, videoService)
	if err != nil {
		return false, err
	}

	// 3.1 Add the first step of the saga, which is to favorite the video
	isFavorited := checkIsFavorited(origin.Status)
	wrapSagaAdd(
		saga,
		&sagaMember{
			method:           addr + videointeractv1.VideoInteract_VideoFavorite_FullMethodName,
			compensateMethod: addr + videointeractv1.VideoInteract_VideoFavoriteRevert_FullMethodName,
			payload: &videointeractv1.VideoFavoriteReq{
				VideoId:    videoId,
				FavoriteId: favoriteId,
				Favorite:   -isFavorited,
			},
		},
		// 3.2 Update User-Video status
		&sagaMember{
			addr + videointeractv1.VideoInteract_SetUserVideoStatus_FullMethodName,
			addr + videointeractv1.VideoInteract_SetUserVideoStatusRevert_FullMethodName,
			&videointeractv1.SetUserVideoStatusReq{
				VideoId:      videoId,
				UserId:       userId,
				OriginStatus: origin.Status,
				Statuses: &videointeractv1.SetUserVideoStatusReq_Statuses{
					IsFavorited: -isFavorited,
				},
			},
		},
	)
	return isFavorited == 1, saga.Submit()
}

func (i *interactRepo) UpvoteVideo(ctx context.Context, userId int64, videoId int64) (bool, error) {
	origin, err := i.data.videoInteractClient.GetUserVideoStatus(ctx, &videointeractv1.GetUserVideoStatusReq{
		VideoId: videoId,
		UserId:  userId,
	})
	if err != nil {
		return false, err
	}

	// Get the video publisher's id.
	resp, err := i.data.videoInfoClient.GetPublisherIdByVideoId(ctx, &videoinfov1.GetPublisherIdByVideoIdReq{
		VideoId: videoId,
	})

	videoAddr, err := resolver.GetRandomAddr(ctx, videoService)
	if err != nil {
		return false, err
	}

	userAddr, err := resolver.GetRandomAddr(ctx, userService)
	if err != nil {
		return false, err
	}
	gid, saga := newSaga(ctx, i.data.dtmServerAddr)
	i.logger.Infof("[UpvoteVideo] gid: %s", gid)

	isUpvoted := checkIsUpvoted(origin.Status)
	wrapSagaAdd(
		saga,
		// 1. Add the first step of the saga, which is to upvote the video
		&sagaMember{
			method:           videoAddr + videointeractv1.VideoInteract_VideoUpvote_FullMethodName,
			compensateMethod: videoAddr + videointeractv1.VideoInteract_VideoUpvoteRevert_FullMethodName,
			payload: &videointeractv1.VideoUpvoteReq{
				VideoId: videoId,
				Upvote:  -isUpvoted, // reverse the status of Upvoted
			},
		},

		// 2. Add the second step of the saga, which is to update the publisher's like count
		&sagaMember{
			method:           userAddr + userinfov1.Userinfo_UpdateUserCntLikes_FullMethodName,
			compensateMethod: userAddr + userinfov1.Userinfo_UpdateUserCntLikesRevert_FullMethodName,
			payload: &userinfov1.UpdateUserCntLikesReq{
				UserId:    resp.PublisherId,
				IsUpvoted: -isUpvoted, // reverse the status of Upvoted
			},
		},

		// 3. Add the third step of the saga, which is to update the user-video status
		&sagaMember{
			method:           videoAddr + videointeractv1.VideoInteract_SetUserVideoStatus_FullMethodName,
			compensateMethod: videoAddr + videointeractv1.VideoInteract_SetUserVideoStatusRevert_FullMethodName,
			payload: &videointeractv1.SetUserVideoStatusReq{
				VideoId:      videoId,
				UserId:       userId,
				OriginStatus: origin.Status,
				Statuses: &videointeractv1.SetUserVideoStatusReq_Statuses{
					IsUpvoted: -isUpvoted,
				},
			},
		},
	)
	return isUpvoted == 1, saga.Submit()
}

func (i *interactRepo) UpvoteBarrage(ctx context.Context, userId int64, barrageId int64) error {
	origin, err := i.data.videoInteractClient.GetUserBarrageStatus(ctx, &videointeractv1.GetUserBarrageStatusReq{
		UserId:    userId,
		BarrageId: barrageId,
	})
	if err != nil {
		return err
	}

	videoAddr, err := resolver.GetRandomAddr(ctx, videoService)
	if err != nil {
		return err
	}

	gid, saga := newSaga(ctx, i.data.dtmServerAddr)
	i.logger.Infof("[UpvoteBarrage] gid: %s", gid)

	isUpvoted := checkIsUpvoted(origin.Status)
	wrapSagaAdd(
		saga,
		// 1. Add the first step of the saga, which is to upvote the video
		&sagaMember{
			method:           videoAddr + videointeractv1.VideoInteract_BarrageUpvote_FullMethodName,
			compensateMethod: videoAddr + videointeractv1.VideoInteract_BarrageUpvoteRevert_FullMethodName,
			payload: &videointeractv1.UpvoteBarrageReq{
				BarrageId: barrageId,
				UserId:    userId,
				Upvote:    -isUpvoted, // reverse the status of Upvoted
			},
		},

		// 2. Add the second step of the saga, which is to update the user-barrage status
		&sagaMember{
			method:           videoAddr + videointeractv1.VideoInteract_SetUserBarrageStatus_FullMethodName,
			compensateMethod: videoAddr + videointeractv1.VideoInteract_SetUserBarrageStatusRevert_FullMethodName,
			payload: &videointeractv1.SetUserBarrageStatusReq{
				UserId:       userId,
				BarrageId:    barrageId,
				OriginStatus: origin.Status,
				Statuses: &videointeractv1.SetUserBarrageStatusReq_Statuses{
					IsUpvoted: -isUpvoted, // reverse the status of Upvoted
				},
			},
		},
	)
	return saga.Submit()
}

func (i *interactRepo) SendBarrage(ctx context.Context, userId int64, req *interactv1.VideoSendBarrageReq) error {
	_, err := i.data.videoInteractClient.VideoSendBarrage(ctx, &videointeractv1.VideoSendBarrageReq{
		PublisherId: userId,
		VideoId:     req.VideoId,
		Content:     req.Content,
		Time:        req.Time,
		Color:       req.Color,
	})
	return err
}

func (i *interactRepo) ShareVideo(ctx context.Context, userId int64, videoId int64) (string, error) {
	// 1. Get the current status of the user-video
	origin, err := i.data.videoInteractClient.GetUserVideoStatus(ctx, &videointeractv1.GetUserVideoStatusReq{
		VideoId: videoId,
		UserId:  userId,
	})
	if err != nil {
		return "", err
	}

	videoAddr, err := resolver.GetRandomAddr(ctx, videoService)
	if err != nil {
		return "", err
	}

	resp, err := i.data.videoInteractClient.VideoShare(ctx, &videointeractv1.VideoShareReq{
		VideoId:  videoId,
		IsShared: checkIsShared(origin.Status),
	})
	if err != nil {
		return "", err
	}
	uri := resp.Uri

	// TODO: 目前先写死返回的url，后续需要改成根据实际情况返回
	url := "http://127.0.0.1:8080" + uri

	isShared := checkIsShared(origin.Status)

	// If isShared == -1, it means the video is not shared by this user before,
	// so we need to add the saga to update the hot count and set the user-video status
	if isShared == -1 {
		gid, saga := newSaga(ctx, i.data.dtmServerAddr)
		i.logger.Infof("[ShareVideo] gid: %s", gid)
		wrapSagaAdd(
			saga,
			&sagaMember{
				videoAddr + videoinfov1.VideoInfo_AddVideoCntShared_FullMethodName,
				videoAddr + videoinfov1.VideoInfo_AddVideoCntSharedRevert_FullMethodName,
				&videoinfov1.AddVideoCntSharedReq{
					VideoId:        videoId,
					IsCompensation: false,
				},
			},
			&sagaMember{
				videoAddr + videointeractv1.VideoInteract_SetUserVideoStatus_FullMethodName,
				videoAddr + videointeractv1.VideoInteract_SetUserVideoStatusRevert_FullMethodName,
				&videointeractv1.SetUserVideoStatusReq{
					UserId:       userId,
					VideoId:      videoId,
					OriginStatus: origin.Status,
					Statuses: &videointeractv1.SetUserVideoStatusReq_Statuses{
						IsShared: -isShared, // reverse the status of Shared
					},
				},
			},
		)
		err = saga.Submit()
		if err != nil {
			return "", err
		}
	}

	return url, nil
}

func (i *interactRepo) ThrowShells(ctx context.Context, userId int64, videoId int64, count int32) error {
	origin, err := i.data.videoInteractClient.GetUserVideoStatus(ctx, &videointeractv1.GetUserVideoStatusReq{
		VideoId: videoId,
		UserId:  userId,
	})
	if err != nil {
		return err
	}
	if origin.ShellsCount+count > 2 {
		return stderr.New("贝壳投递数量已达上限，最多可投递两个☺️")
	}

	resp, err := i.data.videoInfoClient.GetPublisherIdByVideoId(ctx, &videoinfov1.GetPublisherIdByVideoIdReq{
		VideoId: videoId,
	})
	if err != nil {
		return err
	}

	videoAddr, err := resolver.GetRandomAddr(ctx, videoService)
	if err != nil {
		return err
	}

	userAddr, err := resolver.GetRandomAddr(ctx, userService)
	if err != nil {
		return err
	}

	gid, saga := newSaga(ctx, i.data.dtmServerAddr)
	i.logger.Infof("[ThrowShells] gid: %s", gid)

	wrapSagaAdd(
		saga,

		// 1. Update shells of the thrower and the publisher
		&sagaMember{
			method:           userAddr + userinfov1.Userinfo_UpdateUserShells_FullMethodName,
			compensateMethod: userAddr + userinfov1.Userinfo_UpdateUserShellsRevert_FullMethodName,
			payload: &userinfov1.UpdateUserShellsReq{
				UserId:      userId,
				PublisherId: resp.PublisherId,
				Shells:      int64(count),
			},
		},

		// 2. Update the shells the video received
		&sagaMember{
			method:           videoAddr + videointeractv1.VideoInteract_VideoThrowShells_FullMethodName,
			compensateMethod: videoAddr + videointeractv1.VideoInteract_VideoThrowShellsRevert_FullMethodName,
			payload: &videointeractv1.VideoThrowShellsReq{
				VideoId: videoId,
				UserId:  userId,
				Shells:  count,
			},
		},

		// 3. Update the user-video status
		&sagaMember{
			method:           videoAddr + videointeractv1.VideoInteract_SetUserVideoStatus_FullMethodName,
			compensateMethod: videoAddr + videointeractv1.VideoInteract_SetUserVideoStatusRevert_FullMethodName,
			payload: &videointeractv1.SetUserVideoStatusReq{
				VideoId:      videoId,
				UserId:       userId,
				OriginStatus: origin.Status,
				ShellsCount:  origin.ShellsCount + count,
				Statuses: &videointeractv1.SetUserVideoStatusReq_Statuses{
					IsThrownShells: 1,
				},
			}},
	)
	return saga.Submit()
}

func newSaga(ctx context.Context, dtmServerAddr string) (string, *dtmgrpc.SagaGrpc) {
	gid := getid.GetUUID()
	saga := dtmgrpc.NewSagaGrpcWithContext(ctx, dtmServerAddr, gid)
	return gid, saga
}
func wrapSagaAdd(saga *dtmgrpc.SagaGrpc, member ...*sagaMember) *dtmgrpc.SagaGrpc {
	for _, s := range member {
		saga.Add(s.method, s.compensateMethod, s.payload)
	}
	saga.WaitResult = true
	return saga
}
