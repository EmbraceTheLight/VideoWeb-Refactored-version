package userbiz

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
	"vw_gateway/internal/pkg/ecode/errdef/uerr"
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
		return "", helper.HandleError(uerr.ErrFormFileFiled, err)
	}

	err = file.CheckIfPictureValid(fh)
	if err != nil {
		return "", helper.HandleError(uerr.ErrUploadAvatarFailed, err)
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
		return helper.HandleError(uerr.ErrFormFileFiled, err)
	}

	err = file.CheckIfPictureValid(fh)
	if err != nil {
		return helper.HandleError(uerr.ErrUpdateAvatarFailed, err)
	}

	return ufu.repo.UpdateAvatar(ctx, userId, fh)
}
