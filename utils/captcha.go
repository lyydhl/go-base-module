package utils

import (
	"github.com/google/uuid"
	"github.com/mojocn/base64Captcha"
	"image/color"
)

var store = base64Captcha.DefaultMemStore

//configJsonBody json request body.
type configJsonBody struct {
	Id            string
	CaptchaType   string
	VerifyValue   string
	DriverAudio   *base64Captcha.DriverAudio
	DriverString  *base64Captcha.DriverString
	DriverChinese *base64Captcha.DriverChinese
	DriverMath    *base64Captcha.DriverMath
	DriverDigit   *base64Captcha.DriverDigit
}

// 字符串验证码 [数字+字母]
func DriverStringFunc() (id, b64s string, err error) {
	e := configJsonBody{}
	e.Id = uuid.New().String()
	e.DriverString = base64Captcha.NewDriverString(80, 240, 2, 2, 6,
		"23456789abcdefghjkmnpqrstuvwxyz", &color.RGBA{240, 240, 246, 246}, []string{"wqy-microhei.ttc"})
	driver := e.DriverString.ConvertFonts()
	cap := base64Captcha.NewCaptcha(driver, store)
	return cap.Generate()
}

// 纯数字验证码
func DriverDigitFunc() (id, b64s string, err error) {
	e := configJsonBody{}
	e.Id = uuid.New().String()
	e.DriverDigit = base64Captcha.NewDriverDigit(80, 240, 6, 0.7, 80)
	// e.DriverDigit = base64Captcha.DefaultDriverDigit // 默认长度 5
	driver := e.DriverDigit
	cap := base64Captcha.NewCaptcha(driver, store)
	return cap.Generate()
}
