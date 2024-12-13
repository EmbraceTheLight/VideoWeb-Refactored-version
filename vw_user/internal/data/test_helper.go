package data

import (
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"gorm.io/gorm"
	"path"
	"path/filepath"
	"runtime"
	"vw_user/internal/conf"
	"vw_user/internal/data/dal/query"
)

var (
	db *gorm.DB
)

func init() {
	db = t_getGormDB()
	query.SetDefault(db)
}
func getCurrentPath() string {
	_, filename, _, _ := runtime.Caller(1)
	return path.Dir(filename)
}
func t_getGormDB() *gorm.DB {
	c := config.New(
		config.WithSource(
			file.NewSource(filepath.Join(getCurrentPath(), "../../configs/config.yaml")),
		),
	)
	defer c.Close()

	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}
	return NewMySQL(bc.Data)
}

func GetDB() *gorm.DB {
	return db
}
