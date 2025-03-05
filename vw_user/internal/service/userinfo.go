package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	pb "vw_user/api/user/v1/userinfo"
	"vw_user/internal/biz"
)

type UserInfoService struct {
	pb.UnimplementedUserinfoServer
	info   *biz.UserInfoUsecase
	logger *log.Helper
}

func NewUserInfoService(info *biz.UserInfoUsecase, logger log.Logger) *UserInfoService {
	return &UserInfoService{
		info:   info,
		logger: log.NewHelper(logger),
	}
}
func (userInfo *UserInfoService) ForgetPassword(ctx context.Context, req *pb.ForgetPasswordRequest) (rsp *pb.ForgetPasswordResp, err error) {
	return nil, nil
}
func (userInfo *UserInfoService) GetUserDetail(ctx context.Context, req *pb.GetUserDetailRequest) (rsp *pb.GetUserDetailResp, err error) {
	user, level, err := userInfo.info.GetUserDetail(ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	rsp = &pb.GetUserDetailResp{
		Resp: &pb.BaseResp{
			StatusCode: 200,
			Msg:        "success",
		},
		UserDetail: &pb.GetUserDetailResp_UserDetail{
			UserId:     user.UserID,
			Username:   user.Username,
			Email:      user.Email,
			Signature:  user.Signature,
			AvatarPath: user.AvatarPath,
			CntLikes:   user.CntLikes,
			CntFollows: user.CntFollows,
			CntFans:    user.CntFans,
		},
		UserLevel: &pb.GetUserDetailResp_UserLevel{
			NextLevelExp: level.NextExp,
			Exp:          level.Exp,
			Level:        level.Level,
		},
	}
	return rsp, nil
}
func (userInfo *UserInfoService) ModifyEmail(ctx context.Context, req *pb.ModifyEmailRequest) (rsp *pb.ModifyEmailResp, err error) {
	return nil, nil
}
func (userInfo *UserInfoService) ModifyPassword(ctx context.Context, req *pb.ModifyPasswordRequest) (rsp *pb.ModifyPasswordResp, err error) {
	return nil, nil
}
func (userInfo *UserInfoService) ModifySignature(ctx context.Context, req *pb.ModifySignatureRequest) (rsp *pb.ModifySignatureResp, err error) {
	return nil, nil
}
func (userInfo *UserInfoService) ModifyUsername(ctx context.Context, req *pb.ModifyUsernameRequest) (rsp *pb.ModifyUsernameResp, err error) {
	return nil, nil
}
func (userInfo *UserInfoService) UploadAvatar(ctx context.Context, req *pb.UploadAvatarRequest) (rsp *pb.UploadAvatarResp, err error) {
	return nil, nil
}
