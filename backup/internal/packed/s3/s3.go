package s3

import (
	"context"
	"fmt"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"

	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/glog"
)

type S3Client struct {
	endpoint        string
	accessKeyID     string
	secretAccessKey string
	useSSL          bool
}

func New(endpoint, accessKeyID, secretAccessKey string, useSSL bool) *S3Client {
	return &S3Client{
		endpoint:        endpoint,
		accessKeyID:     accessKeyID,
		secretAccessKey: secretAccessKey,
		useSSL:          useSSL,
	}
}

func (client *S3Client) UploadFile(ctx context.Context, bucketName, location, objectName, filePath string) (s3PublicURL string, err error) {
	glog.Debug(ctx, "创建 s3client", client.endpoint, client.accessKeyID, client.secretAccessKey, client.useSSL)
	// Initialize minio client object.
	minioClient, err := minio.New(client.endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(client.accessKeyID, client.secretAccessKey, ""),
		Secure: client.useSSL,
	})

	if err != nil {
		glog.Error(ctx, err)
		return "", err
	}
	glog.Debug(ctx, "创建 s3 连接成功", client.endpoint, ctx)

	// Make a new bucket.
	err = minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: location})

	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := minioClient.BucketExists(ctx, bucketName)
		glog.Debugf(ctx, "桶是否存在 %v，err=\n", exists, errBucketExists)
		if exists {
			glog.Debugf(ctx, "We already own %s\n", bucketName)
		}
	} else {
		glog.Debugf(ctx, "Successfully created %s\n", bucketName)
	}

	// Upload the file with FPutObject
	info, err := minioClient.FPutObject(ctx, bucketName, objectName, filePath, minio.PutObjectOptions{})
	if err != nil {
		glog.Info(ctx, err)
		return "", err
	}

	glog.Infof(ctx, "Successfully uploaded %s of size %d\n", objectName, info.Size)
	return GetPublicURL(ctx, objectName), nil
}

// 获得七牛云访问地址
func GetPublicURL(ctx context.Context, filename string) string {
	config, err := gcfg.New()
	if err != nil {
		glog.Error(ctx, "创建 config 对象失败", err)
	}
	localUrl := config.MustGet(ctx, "storage.s3.accessUrl", false).String()
	fmt.Println("localUrl", localUrl)
	return localUrl + "/" + filename
}
