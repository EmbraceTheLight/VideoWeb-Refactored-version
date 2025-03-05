package captcha

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/spewerspew/spew"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
	"vw_user/internal/conf"
)

func TestImageCaptcha(t *testing.T) {
	id, b64s, ans, err := GenerateGraphicCaptcha()
	require.NoError(t, err)
	spew.Dump(id, b64s, ans)
}

func TestCodeCaptcha(t *testing.T) {
	c := &conf.Email{
		SmtpHost:       "smtp.qq.com",
		SmtpPort:       465,
		SmtpUsername:   "1010642166@qq.com",
		SmtpPassword:   "<PASSWORD>",
		SmtpServername: "smtp.qq.com",
	}
	logger := log.With(log.NewStdLogger(os.Stdout),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
	)
	e := NewEmail(c, logger)
	err := e.SendCode(context.Background(), "10eltzey10@gmail.com", e.CreateVerificationCode())
	require.NoError(t, err)
}
