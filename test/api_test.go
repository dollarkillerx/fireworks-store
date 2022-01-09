package test

import (
	"bytes"
	"fmt"
	"image/color"
	"image/png"
	"io/ioutil"
	"testing"

	"github.com/afocus/captcha"
)

func TestCaptcha(t *testing.T) {
	cap := captcha.New()
	cap.SetFont("comic.ttf")
	// 设置验证码大小
	cap.SetSize(128, 64)
	// 设置干扰强度
	cap.SetDisturbance(captcha.MEDIUM)
	// 设置前景色 可以多个 随机替换文字颜色 默认黑色
	cap.SetFrontColor(color.RGBA{255, 255, 255, 255})
	// 设置背景色 可以多个 随机替换背景色 默认白色
	cap.SetBkgColor(color.RGBA{255, 0, 0, 255}, color.RGBA{0, 0, 255, 255}, color.RGBA{0, 153, 0, 255})
	// 设置字体
	img, str := cap.Create(4, captcha.NUM)
	emptyBuff := bytes.NewBuffer(nil)
	_ = png.Encode(emptyBuff, img)
	ioutil.WriteFile("aaa.png", emptyBuff.Bytes(), 00666)
	fmt.Println(str)
}
