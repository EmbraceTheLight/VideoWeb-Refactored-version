package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"io"
	"mime/multipart"
	"path/filepath"
	"util/helper"
	"vw_gateway/internal/biz"
	"vw_gateway/internal/pkg/ecode/errdef"
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
		return "", helper.HandleError(errdef.ErrUploadAvatarFailed, err)
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
			return "", helper.HandleError(errdef.ErrUploadAvatarFailed, err)
		}

		err = stream.Send(&filev1.UploadAvatarReq{
			Data: &filev1.UploadAvatarReq_FileContent{FileContent: buffer[:n]},
		})
		if err != nil {
			return "", helper.HandleError(errdef.ErrUploadAvatarFailed, err)
		}
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		return "", helper.HandleError(errdef.ErrUploadAvatarFailed, err)
	}
	return res.FilePath, nil
}

func (u *userFileRepo) UpdateAvatar(ctx context.Context, userId int64, fileHeader *multipart.FileHeader) error {
	stream, err := u.data.fileClient.UpdateAvatar(ctx)
	if err != nil {
		return err
	}

	// Send user id and file extension to the USER grpc server, to find the avatar file by it.
	err = stream.Send(&filev1.UpdateAvatarReq{
		Data: &filev1.UpdateAvatarReq_MetaData{
			MetaData: &filev1.UpdateAvatarReq_FileMetadata{
				UserId:        userId,
				FileExtension: filepath.Ext(fileHeader.Filename),
			},
		},
	})
	if err != nil {
		return err
	}

	file, err := fileHeader.Open()
	if err != nil {
		return helper.HandleError(errdef.ErrUpdateFileFailed, err)
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
			return helper.HandleError(errdef.ErrUpdateFileFailed, err)
		}

		err = stream.Send(&filev1.UpdateAvatarReq{
			Data: &filev1.UpdateAvatarReq_FileContent{FileContent: buffer[:n]},
		})
		if err != nil {
			return helper.HandleError(errdef.ErrUpdateFileFailed, err)
		}
	}

	_, err = stream.CloseAndRecv()
	if err != nil {
		return helper.HandleError(errdef.ErrUpdateFileFailed, err)
	}
	return nil
}
