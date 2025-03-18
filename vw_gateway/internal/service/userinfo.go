package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	infov1 "vw_gateway/api/v1/userinfo"
	"vw_gateway/internal/biz"
)

type UserinfoService struct {
	infov1.UnimplementedUserinfoServer
	info   *biz.UserinfoUsecase
	logger *log.Helper
}

func NewUserinfoService(info *biz.UserinfoUsecase, logger log.Logger) *UserinfoService {
	return &UserinfoService{
		info:   info,
		logger: log.NewHelper(logger),
	}
}
func (uis *UserinfoService) ForgetPassword(ctx context.Context, req *infov1.ForgetPasswordReq) (*infov1.ForgetPasswordResp, error) {
	err := uis.info.ForgetPassword(ctx, req.UserId, req.Email, req.InputCode, req.NewPassword)
	if err != nil {
		return nil, err
	}
	return &infov1.ForgetPasswordResp{
		Common: &infov1.CommonResp{
			Status:  200,
			Message: "用户重置密码成功",
		},
	}, nil
}

func (uis *UserinfoService) ModifyEmail(ctx context.Context, req *infov1.ModifyEmailReq) (*infov1.ModifyEmailResp, error) {
	newEmail, err := uis.info.ModifyEmail(ctx, req.UserId, req.Email, req.InputCode)
	if err != nil {
		return nil, err
	}
	return &infov1.ModifyEmailResp{
		Email: newEmail,
		Common: &infov1.CommonResp{
			Status:  200,
			Message: "用户修改邮箱成功",
		},
	}, nil
}

func (uis *UserinfoService) ModifyPassword(ctx context.Context, req *infov1.ModifyPasswordReq) (*infov1.ModifyPasswordResp, error) {
	err := uis.info.ModifyPassword(ctx, req.UserId, req.OldPassword, req.NewPassword)
	if err != nil {
		return nil, err
	}
	return &infov1.ModifyPasswordResp{
		Common: &infov1.CommonResp{
			Status:  200,
			Message: "用户修改密码成功",
		},
	}, nil
}

func (uis *UserinfoService) ModifyUserSignature(ctx context.Context, req *infov1.ModifySignatureReq) (*infov1.ModifySignatureResp, error) {
	newSignature, err := uis.info.ModifyUserSignature(ctx, req.UserId, req.Signature)
	if err != nil {
		return nil, err
	}
	return &infov1.ModifySignatureResp{
		NewSignature: newSignature,
		Common: &infov1.CommonResp{
			Status:  200,
			Message: "用户修改签名成功",
		},
	}, nil
}

func (uis *UserinfoService) Userinfo(ctx context.Context, req *infov1.UserinfoReq) (*infov1.UserinfoResp, error) {
	resp, err := uis.info.Userinfo(ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	return &infov1.UserinfoResp{
		UserInfo: &infov1.UserInfo{
			UserName:   resp.Username,
			Email:      resp.Email,
			Signature:  resp.Signature,
			Shells:     resp.Shells,
			CntFans:    resp.CntFans,
			CntFollows: resp.CntFollows,
			CntVideos:  resp.CntVideos,
			AvatarPath: resp.AvatarPath,
			Gender:     resp.Gender,
			Birthday:   resp.Birthday,
		},
		Common: &infov1.CommonResp{
			Status:  200,
			Message: "获取用户信息成功",
		},
	}, nil
}

func (uis *UserinfoService) ModifyUsername(ctx context.Context, req *infov1.ModifyUsernameReq) (*infov1.ModifyUsernameResp, error) {
	newUsername, err := uis.info.ModifyUsername(ctx, req.UserId, req.NewUsername)
	if err != nil {
		return nil, err
	}
	return &infov1.ModifyUsernameResp{
		NewUsername: newUsername,
		Common: &infov1.CommonResp{
			Status:  200,
			Message: "修改用户名成功",
		},
	}, nil
}
