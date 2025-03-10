package biz

import (
	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/grpc"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"util/helper"
	"util/snowflake"
	filev1 "vw_user/api/v1/userfile"
	"vw_user/internal/pkg/ecode/errdef"
)

type FileRepo interface {
	//Upload(file *os.File, filename string, path string) error
}

type FileUsecase struct {
	filev1.UnimplementedFileServiceServer
	log *log.Helper
	//repo FileRepo
}

func NewFileUsecase(logger log.Logger) *FileUsecase {
	return &FileUsecase{
		log: log.NewHelper(logger),
		//repo: repo,
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
		/* this defer is used to remove the user dir
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
			userID := snowflake.GetID()

			// 2. compute the new user's directory path, with the user id.
			userDir = filepath.Join(resourcePath, strconv.FormatInt(userID, 10))

			err = helper.CreateDir(userDir, os.ModePerm)
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
			// 1. 写入文件
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
