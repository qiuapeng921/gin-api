package captcha

import (
	"github.com/gin-gonic/gin"
	captcha "github.com/mojocn/base64Captcha"
)

var (
	Store     = captcha.DefaultMemStore
	captchaId = "captcha:loveSong"
)

func NewDriver() *captcha.DriverString {
	driver := new(captcha.DriverString)
	driver.Height = 38
	driver.Width = 129
	driver.NoiseCount = 4
	//driver.ShowLineOptions = captcha.OptionShowSineLine | captcha.OptionShowSlimeLine | captcha.OptionShowHollowLine
	driver.ShowLineOptions = captcha.OptionShowSineLine
	driver.Length = 4
	driver.Source = "1234567890qwertyuipkjhgfdsazxcvbnm"
	driver.Fonts = []string{"wqy-microhei.ttc"}
	return driver
}

// 生成图形验证码
func GenerateCaptcha(ctx *gin.Context) {
	var driver = NewDriver().ConvertFonts()
	c := captcha.NewCaptcha(driver, Store)
	_, content, answer := c.Driver.GenerateIdQuestionAnswer()
	item, _ := c.Driver.DrawCaptcha(content)
	c.Store.Set(captchaId, answer)
	_, _ = item.WriteTo(ctx.Writer)
}

// 验证验证码是否正确
func VerifyCaptcha(ctx *gin.Context, code string) bool {
	if Store.Verify(captchaId, code, true) {
		return true
	}
	return false
}
