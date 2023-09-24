package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/mojocn/base64Captcha"
)

func main() {
	// 创建一个字符验证码配置
	captcha := base64Captcha.NewCaptcha(base64Captcha.DefaultDriverDigit, base64Captcha.DefaultMemStore)
	id, b64s, err := captcha.Generate()
	fmt.Println(id, b64s, err)

	// 保存图片
	base64Image, _ := strings.CutPrefix(b64s, "data:image/png;base64,")

	// 解码Base64图片
	imageData, err := base64.StdEncoding.DecodeString(base64Image)
	if err != nil {
		log.Fatalf("Error decoding base64 image: %v", err)
	}

	// 保存图片为PNG文件
	err = ioutil.WriteFile("output.png", imageData, 0644)
	if err != nil {
		log.Fatalf("Error saving image to file: %v", err)
	}
	var s string
	fmt.Println("请输入验证码")
	fmt.Scanln(&s)
	if captcha.Verify(id, s, true) {
		fmt.Println("验证码正确")
	} else {
		fmt.Println("验证码错误")
	}

}
