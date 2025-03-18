package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/types/known/timestamppb"
	followv1 "vw_user/api/v1/follow"
	"vw_user/internal/biz"
)

type FollowService struct {
	followv1.UnimplementedFollowServer
	followUsecase *biz.FollowUsecase
	log           *log.Helper
}

func NewFollowService(followUsecase *biz.FollowUsecase, logger log.Logger) *FollowService {
	return &FollowService{
		followUsecase: followUsecase,
		log:           log.NewHelper(logger),
	}
}

func (s *FollowService) FollowUser(ctx context.Context, req *followv1.FollowUserReq) (*followv1.FollowUserResp, error) {
	err := s.followUsecase.Follow(ctx, req.UserId, req.FollowerUserId, req.FollowListId)
	if err != nil {
		return nil, err
	}
	return &followv1.FollowUserResp{
		StatusCode: int32(codes.OK),
		Message:    "关注成功",
	}, err
}

func (s *FollowService) UnfollowUser(ctx context.Context, req *followv1.UnfollowUserReq) (*followv1.UnfollowUserResp, error) {
	err := s.followUsecase.Unfollow(ctx, req.UserId, req.FolloweeUserId)
	if err != nil {
		return nil, err
	}
	return &followv1.UnfollowUserResp{
		StatusCode: int32(codes.OK),
		Message:    "取消关注成功",
	}, err
}

func (s *FollowService) GetFolloweesInfo(ctx context.Context, req *followv1.GetFolloweesInfoReq) (*followv1.GetFolloweesInfoResp, error) {
	userInfo, err := s.followUsecase.GetFolloweesInfo(ctx, req.UserId, req.FollowListId, req.PageNum, req.PageSize)
	if err != nil {
		return nil, err
	}

	userSummaries := make([]*followv1.GetFolloweesInfoResp_UserSummary, len(userInfo))
	for i, u := range userInfo {
		userSummaries[i] = &followv1.GetFolloweesInfoResp_UserSummary{
			Username:   u.Username,
			Email:      u.Email,
			Signature:  u.Signature,
			AvatarPath: u.AvatarPath,
			Gender:     u.Gender,
			Birthday:   timestamppb.New(u.Birthday),
		}
	}
	
	return &followv1.GetFolloweesInfoResp{
		StatusCode:    int32(codes.OK),
		Message:       "获取关注列表成功",
		FolloweesInfo: userSummaries,
	}, nil
}
