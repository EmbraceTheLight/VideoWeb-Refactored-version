package helper

import (
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/spewerspew/spew"
	"testing"
	"vw_user/internal/conf"
)

func TestGetConfig(t *testing.T) {
	c := config.New(
		config.WithSource(
			file.NewSource("../../../configs/config.yaml"),
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
	spew.Dump(bc)
}
