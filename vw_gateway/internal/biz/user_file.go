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
	filev1 "vw_gateway/api/v1/userfile"
	"vw_gateway/internal/pkg/ecode/errdef"
)

type UserFileRepo interface {
	UploadAvatar(ctx context.Context, fileName string, file *multipart.FileHeader) (filePath string, err error)
}
type UserFileUsecase struct {
	filev1.UnimplementedFileServiceServer
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
	fh, err := helper.FormFile(httpReq, "avatar")
	if err != nil {
		return "", helper.HandleError(errdef.ErrFormFileFiled, err)
	}

	return ufu.repo.UploadAvatar(ctx, fh.Filename, fh)
}
