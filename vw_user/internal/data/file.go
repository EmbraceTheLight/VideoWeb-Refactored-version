package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"vw_user/internal/biz"
	"vw_user/internal/data/dal/model"
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
	u := getQuery(ctx).User
	userDo := u.WithContext(ctx)
	findUser, err := userDo.Where(u.UserID.Eq(userID)).First()
	if err != nil {
		return err
	}
	userDo.ReplaceDB(userDo.UnderlyingDB().Model(findUser))
	_, err = userDo.
		Where(u.UserID.Eq(userID)).
		Updates(&model.User{
			AvatarPath: filePath,
		})
	return err
}
