package models

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"image/png"
	"os"
	"strconv"
	"time"

	openai "github.com/sashabaranov/go-openai"
)

var painterApiKey string
var painterClient *openai.Client //painter client

func NewPainter() {
	painterClient = openai.NewClient(painterApiKey)
}

//返回值为图片URL  /images/generations 目录下
func painterBackURL(prompt string) string {
	prompt = "Parrot on a skateboard performs a trick, cartoon style, natural light, high detail"
	ctx := context.Background()
	reqUrl := openai.ImageRequest{
		Prompt:         prompt,
		Size:           openai.CreateImageSize256x256,
		ResponseFormat: openai.CreateImageResponseFormatURL,
		N:              1,
	}
	respUrl, err := painterClient.CreateImage(ctx, reqUrl)
	if err != nil {
		fmt.Printf("Image creation error: %v\n", err)
		return err.Error()
	}
	return respUrl.Data[0].URL

}

//返回值为图片base64
func painterBackBase64(prompt string) string {
	ctx := context.Background()
	reqBase64 := openai.ImageRequest{
		Prompt:         "Portrait of a humanoid parrot in a classic costume, high detail, realistic light, unreal engine",
		Size:           openai.CreateImageSize256x256,
		ResponseFormat: openai.CreateImageResponseFormatB64JSON,
		N:              1,
	}

	respBase64, err := painterClient.CreateImage(ctx, reqBase64)
	if err != nil {
		fmt.Printf("Image creation error: %v\n", err)
		return err.Error()
	}
	return respBase64.Data[0].B64JSON
}

func generateTimestamp() string {
	ts := time.Now().UnixNano() / 1000000 // 获取当前时间戳
	tsStr := strconv.FormatInt(ts, 10)    // 将时间戳转换为字符串
	return tsStr[len(tsStr)-6:]           // 截取后6位并返回
}

func base64Toimage(B64JSON string) string {
	imgBytes, err := base64.StdEncoding.DecodeString(B64JSON)
	if err != nil {
		fmt.Printf("Base64 decode error: %v\n", err)
		return err.Error()
	}
	r := bytes.NewReader(imgBytes)
	imgData, err := png.Decode(r)
	if err != nil {
		fmt.Printf("PNG decode error: %v\n", err)
		return err.Error()
	}
	imageName := generateTimestamp() + ".png"
	file, err := os.Create(imageName)
	defer file.Close()
	if err != nil {
		fmt.Printf("File creation error: %v\n", err)
		return err.Error()
	}
	if err := png.Encode(file, imgData); err != nil { //前面创建一张空图片，然后写入数据
		fmt.Printf("PNG encode error: %v\n", err)
		return err.Error()
	}
	response := "The image was saved as" + imageName
	return response
}
