package user

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"time"
	idv1 "vw_gateway/api/v1/user/identity"
	"vw_gateway/internal/biz/userbiz"
)

type UserIdentityService struct {
	idv1.UnimplementedIdentityServer
	logger   *log.Helper
	identity *userbiz.UserIdentityUsecase
	captcha  *userbiz.CaptchaUsecase
	info     *userbiz.UserInfo
}

func NewUserIdentityService(identity *userbiz.UserIdentityUsecase, logger log.Logger) *UserIdentityService {
	return &UserIdentityService{
		logger:   log.NewHelper(logger),
		identity: identity,
	}
}

func (uid *UserIdentityService) Register(ctx context.Context, req *idv1.RegisterRequest) (*idv1.RegisterResp, error) {
	birthday, err := time.Parse("2006-01-02", req.Birthday)

	atoken, rtoken, err := uid.identity.Register(ctx, &userbiz.RegisterInfo{
		Username:       req.Username,
		Password:       req.Password,
		RepeatPassword: req.RepeatPassword,
		Gender:         req.Gender,
		Email:          req.Email,
		Signature:      req.Signature,
		Birthday:       birthday,
		InputCode:      req.InputCode,
	})

	if err != nil {
		return nil, err
	}
	return &idv1.RegisterResp{
		StatusCode: 200,
		Msg:        "register successfully",
		Data: &idv1.RegisterResp_Data{
			AccessToken:  atoken,
			RefreshToken: rtoken,
		},
	}, err
}

func (uid *UserIdentityService) Login(ctx context.Context, req *idv1.LoginRequest) (*idv1.LoginResp, error) {
	atoken, rtoken, err := uid.identity.Login(ctx, req.Username, req.Password)
	resp := &idv1.LoginResp{
		StatusCode: 200,
		Msg:        "login successfully",
		Data:       nil,
	}

	if err != nil {
		resp.Msg = err.Error()
		return nil, err
	}
	resp.Data = &idv1.LoginResp_Data{
		AccessToken:  atoken,
		RefreshToken: rtoken,
	}
	return resp, nil
}

func (uid *UserIdentityService) Logout(ctx context.Context, req *idv1.LogoutRequest) (*idv1.LogoutResp, error) {
	err := uid.identity.Logout(ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	return &idv1.LogoutResp{
		StatusCode: 200,
		Msg:        "logout successfully",
	}, nil
}
