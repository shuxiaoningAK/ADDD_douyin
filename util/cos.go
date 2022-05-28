package util

import (
	"context"
	"github.com/tencentyun/cos-go-sdk-v5"
	"mime/multipart"
	"net/http"
	"path"
	"time"
)

var client *cos.Client

func InitCos(url *cos.BaseURL, id, key string) {
	client = cos.NewClient(url, &http.Client{
		//设置超时时间
		Timeout: 100 * time.Second,
		Transport: &cos.AuthorizationTransport{
			//如实填写账号和密钥，也可以设置为环境变量
			SecretID:  id,
			SecretKey: key,
		},
	})
}

func UploadVideo(name string, data *multipart.FileHeader) (string, error) {
	src, err := data.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	name = name + path.Ext(data.Filename)

	_, err = client.Object.Put(context.Background(), name, src, nil)
	if err != nil {
		return "", err
	}
	return client.Object.GetObjectURL(name).String(), nil
}
