package logic

import (
	"context"
	"fmt"
	"io"
	"login-demo/internal/packed/s3"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/glog"
)

var File FileLogic

type FileLogic struct {
}

func (*FileLogic) FileUpload(ctx context.Context, file *ghttp.UploadFile) (err error, filename, fileUrl string) {
	if file == nil {
		return gerror.NewCode(gcode.CodeMissingParameter, "请选择需要上传的文件"), "", ""
	}
	// 效验图片格式，大小
	fileExtension1 := filepath.Ext(file.Filename)
	if fileExtension1 != ".jpg" && fileExtension1 != ".png" && fileExtension1 != ".jpeg" {
		return gerror.NewCode(gcode.CodeMissingParameter, "图片格式不正确，只支持 jpg，png，jpeg 格式图片"), "", ""
	}
	if file.Size > 2*1024*1024 {
		return gerror.NewCode(gcode.CodeMissingParameter, "图片太大，不可超过 2mb"), "", ""
	}

	// 1. 保存到本地目录
	// 生成当前时间的年月日时分秒
	now := time.Now()
	year, month, day := now.Date()
	hour, min, sec := now.Clock()

	// 生成6位随机数
	rand.Seed(time.Now().UnixNano())
	randomNum := rand.Intn(900000) + 100000

	// 构建格式化字符串
	formattedString := fmt.Sprintf("%d%02d%02d%02d%02d%02d-%06d", year, month, day, hour, min, sec, randomNum)

	fileExtension := filepath.Ext(file.Filename)
	file.Filename = formattedString + fileExtension
	config, err := gcfg.New()
	localPath := config.MustGet(ctx, "storage.local.path").String()

	newFilename, err := file.Save(localPath)
	if err != nil {
		return gerror.Wrap(err, "文件保存失败"), "", ""
	}

	filepath := localPath + "/" + newFilename

	glog.Info(ctx, "本地存储地址", filepath)

	// 检查源文件是否存在
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		glog.Infof(ctx, "源文件 %s 不存在\n", filepath)
		return gerror.NewCode(gcode.CodeMissingParameter, "源文件不存在"), "", ""
	}

	// 检查源文件是否可读
	if !isFileReadable(filepath) {
		glog.Infof(ctx, "源文件 %s 不可读\n", filepath)
		return
	}

	// 本地访问路径
	localUrl := config.MustGet(ctx, "storage.local.domain").String()
	filename = newFilename
	fileUrl = localUrl + "/" + newFilename

	glog.Debugf(ctx, "本地访问地址", fileUrl)

	if config.MustGet(ctx, "storage.s3.enabled", false).Bool() {
		endpoint := config.MustGet(ctx, "storage.s3.endpoint").String()
		accessKeyID := config.MustGet(ctx, "storage.s3.accessKeyID").String()
		secretAccessKey := config.MustGet(ctx, "storage.s3.secretAccessKey").String()
		useSSL := config.MustGet(ctx, "storage.s3.useSSL").Bool()
		s3client := s3.New(endpoint, accessKeyID, secretAccessKey, useSSL)

		bucketName := config.MustGet(ctx, "storage.s3.bucketName").String()
		location := config.MustGet(ctx, "storage.s3.location").String()
		glog.Debugf(ctx, "上传七牛云 配置：endpoint: %s,accessKeyID: %s,secretAccessKey: %s, useSSL: %v,bucketName:%s,location:%s\n", endpoint, accessKeyID, secretAccessKey, useSSL, bucketName, location)
		s3PublicURL, err := s3client.UploadFile(ctx, bucketName, location, "images/"+newFilename, filepath)
		if err != nil {
			return gerror.Wrap(err, "文件上传文件服务器失败"), "", ""
		}
		fmt.Println("s3publicUrl", s3PublicURL)
		fileUrl = s3PublicURL
	}
	return
}

// 检查文件是否可读
func isFileReadable(filePath string) bool {
	file, err := os.Open(filePath)
	if err != nil {
		return false
	}
	defer file.Close()
	return true
}

// 复制文件
func copyFile(sourcePath, destinationPath string) error {
	sourceFile, err := os.Open(sourcePath)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destinationFile, err := os.Create(destinationPath)
	if err != nil {
		return err
	}
	defer destinationFile.Close()

	_, err = io.Copy(destinationFile, sourceFile)
	if err != nil {
		return err
	}

	return nil
}
