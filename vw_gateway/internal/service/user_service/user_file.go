package user

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/emptypb"
	"net/http"
	filev1 "vw_gateway/api/v1/user/userfile"
	"vw_gateway/internal/biz/userbiz"
)

type UserFileService struct {
	filev1.UnimplementedFileServiceServer
	userFile *userbiz.UserFileUsecase
	log      *log.Helper
}

func NewUserFileService(userFile *userbiz.UserFileUsecase, logger log.Logger) *UserFileService {
	return &UserFileService{
		userFile: userFile,
		log:      log.NewHelper(logger),
	}
}

func (ufs *UserFileService) UploadAvatar(ctx context.Context, _ *emptypb.Empty) (*filev1.UploadAvatarResp, error) {
	path, err := ufs.userFile.UploadAvatar(ctx)
	if err != nil {
		return nil, err
	}

	return &filev1.UploadAvatarResp{
		FilePath: path,
	}, nil
}

func (ufs *UserFileService) UpdateAvatar(ctx context.Context, req *filev1.UpdateAvatarReq) (*filev1.UpdateAvatarResp, error) {
	err := ufs.userFile.UpdateAvatar(ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	return &filev1.UpdateAvatarResp{
		StatusCode: http.StatusOK,
		Message:    "修改头像成功",
	}, err
}
