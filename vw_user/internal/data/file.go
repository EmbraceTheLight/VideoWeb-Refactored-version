package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
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

func (r *fileRepo) UpdateAvatarPath(ctx context.Context, userID int64, filePath string) error {
	user := getQuery(ctx).User
	userDo := user.WithContext(ctx)
	u, err := userDo.Where(user.UserID.Eq(userID)).First()
	if err != nil {
		return err
	}
	userDo.ReplaceDB(userDo.UnderlyingDB().Model(u))
	_, err = userDo.Where(user.UserID.Eq(userID)).Update(user.AvatarPath, filePath)
	return err
}
