package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	filev1 "vw_user/api/v1/userfile"
	"vw_user/internal/biz"
)

type FileService struct {
	filev1.UnimplementedFileServiceServer
	log         *log.Helper
	fileUsecase *biz.FileUsecase
}

func NewFileService(logger log.Logger, fileUsecase *biz.FileUsecase) *FileService {
	return &FileService{
		log:         log.NewHelper(logger),
		fileUsecase: fileUsecase,
	}
}

func (fs *FileService) UploadAvatar(stream grpc.ClientStreamingServer[filev1.UploadAvatarReq, filev1.UploadAvatarResp]) error {
	return fs.fileUsecase.UploadAvatar(stream)
}

func (fs *FileService) UpdateAvatar(stream grpc.ClientStreamingServer[filev1.UpdateAvatarReq, emptypb.Empty]) error {
	return fs.fileUsecase.UpdateAvatar(stream)
}
