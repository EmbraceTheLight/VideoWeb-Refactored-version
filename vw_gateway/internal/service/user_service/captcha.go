package user

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"vw_gateway/api/v1/user/captcha"
	"vw_gateway/internal/biz/userbiz"
	"vw_gateway/internal/pkg/captcha"
)

type CaptchaService struct {
	captv1.UnimplementedCaptchaServer
	logger  *log.Helper
	captcha *userbiz.CaptchaUsecase
}

func NewCaptchaService(logger log.Logger, captcha *userbiz.CaptchaUsecase) *CaptchaService {
	return &CaptchaService{
		logger:  log.NewHelper(logger),
		captcha: captcha,
	}
}

func (cs *CaptchaService) GetCodeCaptcha(ctx context.Context, req *captv1.GetCodeCaptchaReq) (*captv1.GetCodeCaptchaResp, error) {
	code, err := cs.captcha.SendCodeCaptcha(ctx, req.Email)
	if err != nil {
		return nil, err
	}
	return &captv1.GetCodeCaptchaResp{
		StatusCode: 200,
		Msg:        "get code captcha success",
		Code:       code,
	}, nil
}

func (cs *CaptchaService) GetImageCaptcha(ctx context.Context, req *captv1.GetImageCaptchaRequest) (*captv1.GetImageCaptchaResp, error) {
	id, b64s, ans, err := captcha.GenerateGraphicCaptcha()
	if err != nil {
		return nil, err
	}
	return &captv1.GetImageCaptchaResp{
		StatusCode: 200,
		Msg:        "get image captcha success",
		CaptchaResult: &captv1.GetImageCaptchaResp_CaptchaResult{
			Id:     id,
			B64Log: b64s,
			Answer: ans,
		},
	}, nil
}
