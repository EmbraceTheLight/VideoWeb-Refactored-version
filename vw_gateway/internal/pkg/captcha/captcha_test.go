package captcha_test

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/spewerspew/spew"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
	"vw_gateway/internal/conf"
	"vw_gateway/internal/pkg/captcha"
)

func TestImageCaptcha(t *testing.T) {
	id, b64s, ans, err := captcha.GenerateGraphicCaptcha()
	require.NoError(t, err)
	spew.Dump(id, b64s, ans)
}

func TestCodeCaptcha(t *testing.T) {
	c := &conf.Email{
		SmtpHost:       "<host> e.g. smtp.qq.com",
		SmtpPort:       <Port>,
		SmtpUsername:   "<Email>",
		SmtpPassword:   "<PASSWORD>",
		SmtpServername: "smtp.qq.com",
	}
	logger := log.With(log.NewStdLogger(os.Stdout),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
	)
	e := captcha.NewEmail(c, logger)
	err := e.SendCode(context.Background(), "10eltzey10@gmail.com", e.CreateVerificationCode())
	require.NoError(t, err)
}
