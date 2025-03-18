package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"net/http"
	followv1 "vw_gateway/api/v1/follow"
	"vw_gateway/internal/biz"
)

type FollowService struct {
	followv1.UnimplementedFollowServer
	log           *log.Helper
	followUsecase *biz.FollowUsecase
}

func NewFollowService(logger log.Logger, followUsecase *biz.FollowUsecase) *FollowService {
	return &FollowService{
		log:           log.NewHelper(logger),
		followUsecase: followUsecase,
	}
}

func (fs *FollowService) FollowUser(ctx context.Context, req *followv1.FollowUserReq) (*followv1.FollowUserResp, error) {
	err := fs.followUsecase.FollowOtherUser(ctx, req.UserId, req.FollowerUserId, req.FollowListId)
	if err != nil {
		return nil, err
	}
	return &followv1.FollowUserResp{
		StatusCode: http.StatusOK,
		Message:    "关注用户成功",
	}, nil
}

func (fs *FollowService) UnfollowUser(ctx context.Context, req *followv1.UnfollowUserReq) (*followv1.UnfollowUserResp, error) {
	err := fs.followUsecase.UnfollowOtherUser(ctx, req.UserId, req.FollowedUserId)
	if err != nil {
		return nil, err
	}
	return &followv1.UnfollowUserResp{
		StatusCode: http.StatusOK,
		Message:    "取消关注用户成功",
	}, nil
}

func (fs *FollowService) GetFolloweeInfo(ctx context.Context, req *followv1.GetFolloweeInfoReq) (*followv1.GetFolloweeInfoResp, error) {
	tmp, err := fs.followUsecase.GetFolloweeInfo(ctx, req.UserId, req.FollowListId, req.PageNum, req.PageSize)
	if err != nil {
		return nil, err
	}

	followeesInfo := make([]*followv1.GetFolloweeInfoResp_UserSummary, len(tmp))
	for i, v := range tmp {
		followeesInfo[i] = &followv1.GetFolloweeInfoResp_UserSummary{
			Username:   v.Username,
			Signature:  v.Signature,
			Email:      v.Email,
			Gender:     v.Gender,
			AvatarPath: v.AvatarPath,
			Birthday:   v.Birthday,
		}
	}

	return &followv1.GetFolloweeInfoResp{
		StatusCode:   http.StatusOK,
		Message:      "获取关注的用户信息成功",
		FolloweeInfo: followeesInfo,
	}, nil
}
