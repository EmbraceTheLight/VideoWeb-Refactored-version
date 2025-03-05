package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/protobuf/types/known/emptypb"
	"time"
	"vw_user/internal/pkg/ecode/errdef"

	idv1 "vw_user/api/v1/identity"
	"vw_user/internal/biz"
)

type UserIdentityService struct {
	idv1.UnimplementedIdentityServer
	identity *biz.UserIdentityUsecase
	logger   *log.Helper
}

func NewUserIdentityService(identity *biz.UserIdentityUsecase, logger log.Logger) *UserIdentityService {
	return &UserIdentityService{
		identity: identity,
		logger:   log.NewHelper(logger),
	}
}

func (s *UserIdentityService) CheckUsernamePassword(ctx context.Context, req *idv1.CheckUsernamePasswordReq) (*idv1.CheckUsernamePasswordResp, error) {
	username, inputPasswd, correctPasswd := req.Username, req.Password, req.CorrectPassword
	if username == "" || inputPasswd == "" {
		if username == "" {
			return nil, errdef.ErrUserNameEmpty
		} else { //password == ""
			return nil, errdef.ErrUserPasswordEmpty
		}
	}

	// check password
	if err := bcrypt.CompareHashAndPassword([]byte(correctPasswd), []byte(inputPasswd)); err != nil {
		return nil, errdef.ErrUserPasswordError
	}
	return &idv1.CheckUsernamePasswordResp{
		Code:    "200",
		Message: "success",
	}, nil
}

func (s *UserIdentityService) CacheAccessToken(ctx context.Context, req *idv1.CacheAccessTokenReq) (*emptypb.Empty, error) {
	return nil, s.identity.CacheAccessToken(ctx, req.AccessToken, req.Expiration.AsDuration())
}

func (s *UserIdentityService) AddExpForLogin(ctx context.Context, req *idv1.AddExpForLoginReq) (*emptypb.Empty, error) {
	return nil, s.identity.AddExpForLogin(ctx, req.UserId)
}

func (s *UserIdentityService) Register(ctx context.Context, req *idv1.RegisterReq) (*idv1.RegisterResp, error) {
	birthday, _ := time.Parse("2006-01-02", req.Birthday)
	userID, isAdmin, err := s.identity.Register(ctx, &biz.RegisterInfo{
		UserID:         new(int64),
		Username:       req.Username,
		Password:       req.Password,
		RepeatPassword: req.RepeatPassword,
		Gender:         req.Gender,
		InputCode:      req.InputCode,
		VerifyCode:     req.VerifyCode,
		Email:          req.Email,
		Signature:      req.Signature,
		Birthday:       birthday,
	})
	if err != nil {
		return nil, err
	}
	return &idv1.RegisterResp{
		StatusCode: 200,
		Msg:        "user: " + req.Username + " register successfully.",
		UserID:     userID,
		IsAdmin:    isAdmin,
	}, nil
}

func (s *UserIdentityService) Logout(ctx context.Context, req *idv1.LogoutReq) (*emptypb.Empty, error) {
	return nil, s.identity.Logout(ctx, req.AccessToken)
}
