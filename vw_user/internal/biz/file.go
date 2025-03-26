package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"time"
	utilCtx "util/context"
	"util/getid"
	"util/helper"
	"util/helper/file"
	filev1 "vw_user/api/v1/userfile"
	"vw_user/internal/pkg/ecode/errdef"
)

type FileRepo interface {
	UpdateAvatarPath(ctx context.Context, userID int64, filePath string) error
}

type FileUsecase struct {
	log  *log.Helper
	repo FileRepo
}

func NewFileUsecase(logger log.Logger, repo FileRepo) *FileUsecase {
	return &FileUsecase{
		log:  log.NewHelper(logger),
		repo: repo,
	}
}

func (u *FileUsecase) UploadAvatar(stream grpc.ClientStreamingServer[filev1.UploadAvatarReq, filev1.UploadAvatarResp]) error {
	var (
		avatarFilePath string
		userDir        string
		avatarFile     *os.File
		err            error
	)

	defer func() {
		/* this defer is used to remove the userbiz dir
		if there is any error Behind this logic*/
		if avatarFile != nil {
			avatarFile.Close()
		}
		if err != nil {
			_ = os.RemoveAll(userDir)
		}
	}()

	for {
		var req *filev1.UploadAvatarReq
		req, err = stream.Recv()
		if err == io.EOF { // 传输完成
			break
		}
		if err != nil {
			return err
		}

		// Handle the data，there are two types of data：
		//One is avatarFile name(string);
		//the other is avatarFile content([]byte).
		switch data := req.Data.(type) {

		// * Get avatarFile name, then create the avatarFile。
		case *filev1.UploadAvatarReq_FileName:
			// Create User Dir
			// 1. Generate UserID by Snowflake algorithm.
			userID := getid.GetID()

			// 2. compute the new userbiz's directory path, with the userbiz id.
			userDir = filepath.Join(resourcePath, strconv.FormatInt(userID, 10))

			err = file.CreateDir(userDir, os.ModePerm)
			if err != nil {
				return helper.HandleError(errdef.ErrCreateUserDirFailed, err)
			}

			// 3. Create the avatar avatarFile by the given avatarFile name.
			avatarFilePath = filepath.Join(userDir, "avatar"+filepath.Ext(data.FileName))
			avatarFile, err = os.Create(avatarFilePath)
			if err != nil {
				return helper.HandleError(errdef.ErrCreateUserAvatarFailed, err)
			}

		// * Get avatarFile content
		case *filev1.UploadAvatarReq_FileContent:
			// 1. Write data to the avatar file.
			fileContent := data.FileContent
			_, err = avatarFile.Write(fileContent)
			if err != nil {
				return helper.HandleError(errdef.ErrSaveAvatarFailed, err)
			}
		}
	}

	err = stream.SendAndClose(&filev1.UploadAvatarResp{FilePath: avatarFilePath})
	return err
}

func (u *FileUsecase) UpdateAvatar(stream grpc.ClientStreamingServer[filev1.UpdateAvatarReq, emptypb.Empty]) error {
	var (
		avatarFilePath string
		newFileName    string
		avatarFile     *os.File
		userID         int64
		err            error
	)

	// Close the avatar file, if it is not closed normally.
	// if it has been closed normally, discard the error.
	defer func() {
		_ = avatarFile.Close()
	}()

	for {
		var req *filev1.UpdateAvatarReq
		req, err = stream.Recv()
		if err == io.EOF { // 传输完成
			break
		}
		if err != nil {
			return err
		}

		// Handle the data，there are two types of data：
		// One is userbiz id(int64);
		// the other is avatarFile content([]byte).
		switch data := req.Data.(type) {

		// * Get avatarFile name, then create the avatarFile。
		case *filev1.UpdateAvatarReq_MetaData:
			// Delete userbiz's old avatar
			// 1. compute the new userbiz's directory avatarFilePath, with the userbiz id.
			userID = data.MetaData.UserId
			userDir := filepath.Join(resourcePath, strconv.FormatInt(userID, 10))

			// 2. Find and prepare to the old avatar, by O_TRUNC and O_WRONLY flags.
			avatarFilePath, err = file.NewFileSearcher().Find(userDir, avatar)
			if err != nil {
				return helper.HandleError(errdef.ErrUpdateAvatarFailed, err)
			}

			// 3. Open the old avatar avatarFile by the given avatarFile name.
			avatarFile, err = os.OpenFile(avatarFilePath, os.O_WRONLY|os.O_TRUNC, 0666)

			// 4. Get the new avatar avatarFile name. Because the avatar file's extension may be changed.
			newFileName = avatar + data.MetaData.FileExtension

		// * Get avatarFile content
		case *filev1.UpdateAvatarReq_FileContent:
			// Write data to the truncated avatar avatarFile.
			fileContent := data.FileContent
			_, err = avatarFile.Write(fileContent)
			if err != nil {
				return helper.HandleError(errdef.ErrUpdateAvatarFailed, err)
			}
		}
	}

	// Close the avatar file, so that it can be renamed.
	err = avatarFile.Close()
	if err != nil {
		return helper.HandleError(errdef.ErrUpdateAvatarFailed, err)
	}

	// rename the old avatar avatarFile to the new avatar avatarFile name.
	baseDir := filepath.Dir(avatarFilePath)
	err = os.Rename(avatarFilePath, filepath.Join(baseDir, newFileName))
	if err != nil {
		return helper.HandleError(errdef.ErrUpdateAvatarFailed, err)
	}

	// Update the userbiz's avatar path in the database.
	ctx, cancel := utilCtx.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = u.repo.UpdateAvatarPath(ctx, userID, filepath.Join(baseDir, newFileName))
	if err != nil {
		return helper.HandleError(errdef.ErrUpdateAvatarFailed, err)
	}

	err = stream.SendAndClose(nil)
	return err
}
