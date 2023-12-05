package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/aliyun/fc-runtime-go-sdk/fc"
	qrcode "github.com/skip2/go-qrcode"
)

func main() {
	fmt.Println("Hello Serverless World!")

	png, err := qrcode.Encode("opt.Content", qrcode.Medium, 256)
	fmt.Println(png, err)

	fc.StartHttp(HandleGenQRCode)
}

type AliyunOssConf struct {
	Cname        string `json:"cname"`
	EndPoint     string `json:"endpoint"`
	AccessKey    string `json:"accessKey"`
	AccessSecret string `json:"accesssSecret"`
	Bucket       string `json:"bucket"`
}

func HandleGenQRCode(ctx context.Context, w http.ResponseWriter, req *http.Request) error {

	content := req.FormValue("content")
	png, err := qrcode.Encode(content, qrcode.Medium, 256)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Add("Content-Type", "text/plain")
		w.Write([]byte(err.Error()))
		return nil
	}

	fmt.Println(png)
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "text/plain")
	w.Write(png)
	return nil

}
