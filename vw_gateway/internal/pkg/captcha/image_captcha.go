package captcha

import (
	"github.com/mojocn/base64Captcha"
	"time"
)

var (
	expiration = 5 * time.Minute // 5 minutes
	store      = base64Captcha.NewMemoryStore(base64Captcha.GCLimitNumber, expiration)
)

// GenerateGraphicCaptcha generates a captcha image.
func GenerateGraphicCaptcha() (id, b64s, ans string, err error) {
	randRGBA := base64Captcha.RandColor()
	var driverString = base64Captcha.DriverString{
		Height:          75,
		Width:           150,
		NoiseCount:      0,
		ShowLineOptions: base64Captcha.OptionShowHollowLine | base64Captcha.OptionShowSlimeLine,
		Length:          4,
		Source:          "abcdefghijklmnopqrstuvwxyz0123456789",
		BgColor:         &randRGBA,
		Fonts:           []string{"wqy-microhei.ttc"},
	}
	var driver base64Captcha.Driver = driverString.ConvertFonts()
	captcha := base64Captcha.NewCaptcha(driver, store)
	id, b64s, ans, err = captcha.Generate()
	return
}

// GraphicCaptchaVerify verifies the image captcha.
func GraphicCaptchaVerify(id, capt string) bool {
	//if the clear flag of store.Verify is true,which means the image captcha can use only once,
	//else,if the clear flag is false,which means the image captcha can use multiple times
	return store.Verify(id, capt, true)
}
