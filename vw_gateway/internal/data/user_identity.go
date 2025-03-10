package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/durationpb"
	"time"
	"vw_gateway/internal/biz"
	idv1 "vw_user/api/v1/identity"
)

type userIdentityRepo struct {
	logger *log.Helper
	data   *Data
}

func NewUserIdentityRepo(data *Data, logger log.Logger) biz.UserIdentityRepo {
	return &userIdentityRepo{
		data:   data,
		logger: log.NewHelper(logger),
	}
}

func (u *userIdentityRepo) CacheAccessToken(ctx context.Context, accessToken string, expiration time.Duration) error {
	_, err := u.data.userIdentityClient.CacheAccessToken(ctx, &idv1.CacheAccessTokenReq{
		AccessToken: accessToken,
		Expiration:  durationpb.New(expiration),
	})
	return err
}

func (u *userIdentityRepo) Register(ctx context.Context, registerInfo *biz.RegisterInfo) (userID int64, isAdmin bool, err error) {
	resp, err := u.data.userIdentityClient.Register(ctx, &idv1.RegisterReq{
		Username:       registerInfo.Username,
		Password:       registerInfo.Password,
		RepeatPassword: registerInfo.RepeatPassword,
		Gender:         registerInfo.Gender,
		Email:          registerInfo.Email,
		Birthday:       registerInfo.Birthday.Format("2006-01-02"),
		Signature:      registerInfo.Signature,
		InputCode:      registerInfo.InputCode,
	})
	if err != nil {
		return 0, false, err
	}
	return resp.UserID, resp.IsAdmin, err
}

func (u *userIdentityRepo) Logout(ctx context.Context, accessToken string) error {
	_, err := u.data.userIdentityClient.Logout(ctx, &idv1.LogoutReq{
		AccessToken: accessToken,
	})
	return err
}

func (u *userIdentityRepo) Login(ctx context.Context, username, password string) (*biz.LoginResponse, error) {
	resp, err := u.data.userIdentityClient.Login(ctx, &idv1.LoginReq{
		Username: username,
		Password: password,
	})
	if err != nil {
		return nil, err
	}
	return &biz.LoginResponse{
		UserId:   resp.UserId,
		Username: resp.Username,
		IsAdmin:  resp.IsAdmin,
	}, err
}
