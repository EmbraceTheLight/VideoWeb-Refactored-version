package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/durationpb"
	"time"
	"vw_gateway/internal/biz"
	idv1 "vw_user/api/v1/identity"
	infov1 "vw_user/api/v1/userinfo"
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

func (u *userIdentityRepo) GetUserSummaryInfoByUsername(ctx context.Context, username string) (*biz.UserSummaryInfo, error) {
	resp, err := u.data.userInfoClient.GetUserSummaryInfoByUsername(ctx, &infov1.GetUserSummaryInfoByUsernameReq{Username: username})
	if err != nil {
		return nil, err
	}
	return &biz.UserSummaryInfo{
		UserID:     resp.UserId,
		Username:   resp.Username,
		Password:   resp.Password,
		Email:      resp.Email,
		Gender:     int(resp.Gender),
		Signature:  resp.Signature,
		AvatarPath: resp.AvatarPath,
		Birthday:   resp.Birthday.AsTime(),
		IsAdmin:    resp.IsAdmin,
	}, nil
}

func (u *userIdentityRepo) CheckUsernamePassword(ctx context.Context, username, inputPassword, correctPassword string) error {
	_, err := u.data.userIdentityClient.CheckUsernamePassword(ctx, &idv1.CheckUsernamePasswordReq{Username: username, Password: inputPassword, CorrectPassword: correctPassword})
	if err != nil {
		return err
	}
	return nil
}

func (u *userIdentityRepo) CacheAccessToken(ctx context.Context, accessToken string, expiration time.Duration) error {
	_, err := u.data.userIdentityClient.CacheAccessToken(ctx, &idv1.CacheAccessTokenReq{
		AccessToken: accessToken,
		Expiration:  durationpb.New(expiration),
	})
	return err
}

func (u *userIdentityRepo) AddExpForLogin(ctx context.Context, userID int64) error {
	_, err := u.data.userIdentityClient.AddExpForLogin(ctx, &idv1.AddExpForLoginReq{
		UserId: userID,
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
		VerifyCode:     registerInfo.VerifyCode,
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
