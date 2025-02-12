TEST DEMO FILE. PLEASE EDIT THE NECESSARY DATA.

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

func TestGenerateImageCaptcha(t *testing.T) {
	id, b64s, ans, err := GenerateGraphicCaptcha()
	require.NoError(t, err)
	spew.Dump(id, b64s, ans)
}

func TestGenerateCodeCaptcha(t *testing.T) {
	/******************EDIT THE FOLLOWING PART************************/
	c := &conf.Email{
			SmtpHost:       "smtp.qq.com",
			SmtpPort:       465,
			SmtpUsername:   "<Your email address>",
			SmtpPassword:   "<Your email authorization code>",
			SmtpServername: "smtp.qq.com",
	}
	/*****************************************************************/

	logger := log.With(log.NewStdLogger(os.Stdout),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
	)
	e := NewEmail(c, logger)

	/******************EDIT THE FOLLOWING PART************************/
	err := e.SendCode(context.Background(), "<Target email address>", e.CreateVerificationCode())
	/*****************************************************************/

	require.NoError(t, err)
}
