package data

import (
	"context"
	"fmt"
	kerr "github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"io"
	"mime/multipart"
	"vw_gateway/internal/biz"
	filev1 "vw_user/api/v1/userfile"
)

const (
	kb = 1024
	mb = 1024 * kb
)

type userFileRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserFileRepo(data *Data, logger log.Logger) biz.UserFileRepo {
	return &userFileRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "user_file/repo")),
	}
}

func (u *userFileRepo) UploadAvatar(ctx context.Context, fileName string, fileHeader *multipart.FileHeader) (filePath string, err error) {
	stream, err := u.data.fileClient.UploadAvatar(ctx)
	if err != nil {
		return "", err
	}

	err = stream.Send(&filev1.UploadAvatarReq{
		Data: &filev1.UploadAvatarReq_FileName{
			FileName: fileName,
		},
	})
	if err != nil {
		return "", err
	}

	file, err := fileHeader.Open()

	if err != nil {
		return "", kerr.New(50001, "open file failed: "+err.Error(), "上传头像失败")
	}
	defer file.Close()

	// Transfer file content to the USER grpc server
	buffer := make([]byte, 32*kb)
	for {
		n, err := file.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			return "", kerr.New(50001, "open file failed: "+err.Error(), "上传头像失败")
		}

		err = stream.Send(&filev1.UploadAvatarReq{
			Data: &filev1.UploadAvatarReq_FileContent{FileContent: buffer[:n]},
		})
		if err != nil {
			return "", kerr.New(50001, "open file failed: "+err.Error(), "上传头像失败")
		}
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		return "", fmt.Errorf("failed to receive response: %w", err)
	}
	return res.FilePath, nil
}
