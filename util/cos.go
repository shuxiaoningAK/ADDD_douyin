package util

import (
	"ADDD_DOUYIN/model"
	"context"
	"github.com/tencentyun/cos-go-sdk-v5"
	"mime/multipart"
	"net/http"
	"path"
	"time"
)

var client *cos.Client
var baseUrl string

func InitCos(url *cos.BaseURL, id, key string) {
	baseUrl = url.BucketURL.String()
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

func UploadVideo(name string, data *multipart.FileHeader, video *model.Video) error {
	src, err := data.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	withVideoExt := name + path.Ext(data.Filename)
	withImgExt := name + "_0.jpg"

	if _, err = client.Object.Put(context.Background(), withVideoExt, src, nil); err != nil {
		return err
	} else {
		video.PlayUrl = baseUrl + "/" + withVideoExt
		video.CoverUrl = baseUrl + "/" + withImgExt
		return nil
	}

}
