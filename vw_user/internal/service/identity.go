package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	khttp "github.com/go-kratos/kratos/v2/transport/http"
	"time"
	pb "vw_user/api/user/v1/identity"
	"vw_user/internal/biz"
	"vw_user/internal/pkg/captcha"
	"vw_user/internal/pkg/ecode/errdef"
)

type UserIdentityService struct {
	pb.UnimplementedIdentityServer
	pb.UnimplementedCaptchaServer
	identity *biz.UserIdentityUsecase
	logger   *log.Helper
}

func NewUserIdentityService(identity *biz.UserIdentityUsecase, logger log.Logger) *UserIdentityService {
	return &UserIdentityService{
		identity: identity,
		logger:   log.NewHelper(logger),
	}
}

func (s *UserIdentityService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResp, error) {
	username := req.Username
	password := req.Password
	if username == "" || password == "" {
		if username == "" {
			return nil, errdef.ErrUserNameEmpty
		} else { //password == ""
			return nil, errdef.ErrUserPasswordEmpty
		}
	}
	atoken, rtoken, err := s.identity.Login(ctx, username, password)
	if err != nil {
		return nil, err
	}
	return &pb.LoginResp{
		StatusCode: 200,
		Msg:        "login success",
		Data: &pb.LoginResp_Data{
			AccessToken:  atoken,
			RefreshToken: rtoken,
		},
	}, nil
}

func (s *UserIdentityService) Logout(ctx context.Context, req *pb.LogoutRequest) (*pb.LogoutResp, error) {
	return nil, nil
}
func (s *UserIdentityService) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResp, error) {
	//birthday is provided by front-end in format "2006-01-02"，we need not check if time.Parse return an error
	birthday, _ := time.Parse("2006-01-02", req.Birthday)
	httpCtx, ok := khttp.RequestFromServerContext(ctx)
	if !ok {
		return nil, errors.New(500, "invalid context", "服务器内部错误")
	}
	atoken, rtoken, err := s.identity.Register(ctx, httpCtx, &biz.RegisterInfo{
		Username:       req.Username,
		Password:       req.Password,
		RepeatPassword: req.RepeatPassword,
		Email:          req.Email,
		Gender:         req.Gender,
		Signature:      req.Signature,
		Code:           req.Code,
		Birthday:       birthday,
	})
	if err != nil {
		return nil, err
	}
	return &pb.RegisterResp{
		StatusCode: 200,
		Msg:        "register success",
		Data: &pb.RegisterResp_Data{
			AccessToken:  atoken,
			RefreshToken: rtoken,
		},
	}, nil
}

func (s *UserIdentityService) GetCodeCaptcha(ctx context.Context, req *pb.GetCodeCaptchaRequest) (*pb.GetCodeCaptchaResp, error) {
	code, err := s.identity.SendCodeCaptcha(ctx, req.Email)
	if err != nil {
		return nil, err
	}
	return &pb.GetCodeCaptchaResp{
		StatusCode: 200,
		Msg:        "get code captcha success",
		Code:       code,
	}, nil
}
func (s *UserIdentityService) GetImageCaptcha(ctx context.Context, req *pb.GetImageCaptchaRequest) (*pb.GetImageCaptchaResp, error) {
	id, b64s, ans, err := captcha.GenerateGraphicCaptcha()
	if err != nil {
		return nil, err
	}
	return &pb.GetImageCaptchaResp{
		StatusCode: 200,
		Msg:        "get image captcha success",
		CaptchaResult: &pb.GetImageCaptchaResp_CaptchaResult{
			Id:     id,
			B64Log: b64s,
			Answer: ans,
		},
	}, nil
}
