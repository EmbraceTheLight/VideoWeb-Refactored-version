package biz

import (
	"context"
	stderr "errors"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	khttp "github.com/go-kratos/kratos/v2/transport/http"
	"mime/multipart"
	"net/http"
	"util/helper"
	"util/helper/file"
	"vw_gateway/internal/pkg/ecode/errdef"
)

type UserFileRepo interface {
	UploadAvatar(ctx context.Context, fileName string, file *multipart.FileHeader) (filePath string, err error)
	UpdateAvatar(ctx context.Context, userId int64, file *multipart.FileHeader) error
}
type UserFileUsecase struct {
	repo UserFileRepo
	log  *log.Helper
}

func NewUserFileUsecase(repo UserFileRepo, logger log.Logger) *UserFileUsecase {
	return &UserFileUsecase{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}

func (ufu *UserFileUsecase) UploadAvatar(ctx context.Context) (filePath string, err error) {
	// Parse form data and get avatar file
	httpReq, ok := khttp.RequestFromServerContext(ctx)
	if !ok {
		return "", helper.HandleError(
			errors.New(http.StatusInternalServerError, "invalid context", "服务器内部错误"),
			stderr.New("khttp.RequestFromServerContext(ctx) failed"))
	}
	fh, err := file.FormFile(httpReq, "avatar")
	if err != nil {
		return "", helper.HandleError(errdef.ErrFormFileFiled, err)
	}

	err = file.CheckPictureValid(fh)
	if err != nil {
		return "", helper.HandleError(errdef.ErrUploadAvatarFailed, err)
	}

	return ufu.repo.UploadAvatar(ctx, fh.Filename, fh)
}

func (ufu *UserFileUsecase) UpdateAvatar(ctx context.Context, userId int64) error {
	// Parse form data and get avatar file
	httpReq, ok := khttp.RequestFromServerContext(ctx)
	if !ok {
		return helper.HandleError(
			errors.New(http.StatusInternalServerError, "invalid context", "服务器内部错误"),
			stderr.New("khttp.RequestFromServerContext(ctx) failed"))
	}
	fh, err := file.FormFile(httpReq, "avatar")
	if err != nil {
		return helper.HandleError(errdef.ErrFormFileFiled, err)
	}

	err = file.CheckPictureValid(fh)
	if err != nil {
		return helper.HandleError(errdef.ErrUpdateFileFailed, err)
	}

	return ufu.repo.UpdateAvatar(ctx, userId, fh)
}
