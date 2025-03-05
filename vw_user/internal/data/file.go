package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"os"
	"vw_user/internal/biz"
)

type fileRepo struct {
	data *Data
	log  *log.Helper
}

func NewFileRepo(data *Data, logger log.Logger) biz.FileRepo {
	return &fileRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (f *fileRepo) Upload(file *os.File, filename string, path string) error {

	//TODO implement me
	panic("implement me")
}
