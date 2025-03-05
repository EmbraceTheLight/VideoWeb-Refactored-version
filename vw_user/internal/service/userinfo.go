package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/timestamppb"
	infov1 "vw_user/api/v1/userinfo"
	"vw_user/internal/biz"
)

type UserInfoService struct {
	infov1.UnimplementedUserinfoServer
	info   *biz.UserInfoUsecase
	logger *log.Helper
}

func NewUserInfoService(info *biz.UserInfoUsecase, logger log.Logger) *UserInfoService {
	return &UserInfoService{
		info:   info,
		logger: log.NewHelper(logger),
	}
}

func (userInfo *UserInfoService) GetUserSummaryInfoByUsername(ctx context.Context, req *infov1.GetUserSummaryInfoByUsernameReq) (*infov1.GetUserSummaryInfoByUsernameResp, error) {
	ret, err := userInfo.info.GetUserSummaryInfoByUsername(ctx, req.Username)
	if err != nil {
		return nil, err
	}

	return &infov1.GetUserSummaryInfoByUsernameResp{
		UserId:     ret.UserID,
		Username:   ret.Username,
		Password:   ret.Password,
		Email:      ret.Email,
		Gender:     int32(ret.Gender),
		Signature:  ret.Signature,
		AvatarPath: ret.AvatarPath,
		Birthday:   timestamppb.New(ret.Birthday),
		IsAdmin:    ret.IsAdmin,
	}, nil
}

func (userInfo *UserInfoService) ForgetPassword(ctx context.Context, req *infov1.ForgetPasswordRequest) (rsp *infov1.ForgetPasswordResp, err error) {
	return nil, nil
}
func (userInfo *UserInfoService) GetUserDetail(ctx context.Context, req *infov1.GetUserDetailRequest) (rsp *infov1.GetUserDetailResp, err error) {
	user, level, err := userInfo.info.GetUserDetail(ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	rsp = &infov1.GetUserDetailResp{
		Resp: &infov1.BaseResp{
			StatusCode: 200,
			Msg:        "success",
		},
		UserDetail: &infov1.GetUserDetailResp_UserDetail{
			UserId:     user.UserID,
			Username:   user.Username,
			Email:      user.Email,
			Signature:  user.Signature,
			AvatarPath: user.AvatarPath,
			CntLikes:   user.CntLikes,
			CntFollows: user.CntFollows,
			CntFans:    user.CntFans,
		},
		UserLevel: &infov1.GetUserDetailResp_UserLevel{
			NextLevelExp: level.NextExp,
			Exp:          level.Exp,
			Level:        level.Level,
		},
	}
	return rsp, nil
}
func (userInfo *UserInfoService) ModifyEmail(ctx context.Context, req *infov1.ModifyEmailRequest) (rsp *infov1.ModifyEmailResp, err error) {
	return nil, nil
}
func (userInfo *UserInfoService) ModifyPassword(ctx context.Context, req *infov1.ModifyPasswordRequest) (rsp *infov1.ModifyPasswordResp, err error) {
	return nil, nil
}
func (userInfo *UserInfoService) ModifySignature(ctx context.Context, req *infov1.ModifySignatureRequest) (rsp *infov1.ModifySignatureResp, err error) {
	return nil, nil
}
func (userInfo *UserInfoService) ModifyUsername(ctx context.Context, req *infov1.ModifyUsernameRequest) (rsp *infov1.ModifyUsernameResp, err error) {
	return nil, nil
}
func (userInfo *UserInfoService) UploadAvatar(ctx context.Context, req *infov1.UploadAvatarRequest) (rsp *infov1.UploadAvatarResp, err error) {
	return nil, nil
}
