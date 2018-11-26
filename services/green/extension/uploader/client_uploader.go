package green_extension_uploader

import (
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/green"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/satori/go.uuid"
	"os"
	"bytes"
	"time"
	"encoding/json"
)

type uploadCredentials struct {
	AccessKeyId string `json:"accessKeyId"`
	AccessKeySecret string `json:"accessKeySecret"`
	SecurityToken string `json:"securityToken"`
	ExpiredTime int64 `json:"expiredTime"`
	OssEndpoint string `json:"ossEndpoint"`
	UploadBucket string `json:"uploadBucket"`
	UploadFolder string `json:"uploadFolder"`
}


type ClientUploader struct {
	GreenClient *green.Client
	UploadCredentials *uploadCredentials
	Headers map[string]string
	Prefix string
}



func GetImageUploaderClient(client *green.Client)(clientUploader *ClientUploader)  {
	clientUploader = &ClientUploader{client, nil, make(map[string]string), "images"}
	return
}

func GetVideoUploaderClient(client *green.Client)(clientUploader *ClientUploader)  {
	clientUploader = &ClientUploader{client, nil, make(map[string]string), "videos"}
	return
}

func GetVoiceUploaderClient(client *green.Client)(clientUploader *ClientUploader)  {
	clientUploader = &ClientUploader{client, nil, make(map[string]string), "voices"}
	return
}

func GetFileUploaderClient(client *green.Client)(clientUploader *ClientUploader)  {
	clientUploader = &ClientUploader{client, nil, make(map[string]string), "files"}
	return
}


func (clientUploader *ClientUploader)UploadFile(filepath string)(*string, error)  {
	uploadCredentials, err := clientUploader.getUploadCredentials()
	fmt.Println(uploadCredentials.OssEndpoint)
	if uploadCredentials == nil {
		return nil, err
	}

	client, err := oss.New(uploadCredentials.OssEndpoint, uploadCredentials.AccessKeyId, uploadCredentials.AccessKeySecret, oss.SecurityToken(uploadCredentials.SecurityToken))
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}
	// 获取存储空间。
	bucket, err := client.Bucket(uploadCredentials.UploadBucket)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	// 读取本地文件。
	fd, err := os.Open(filepath)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}
	defer fd.Close()

	uuid := uuid.NewV4()
	object := uploadCredentials.UploadFolder + "/" + clientUploader.Prefix + "/" + uuid.String()
	// 上传文件流。
	err = bucket.PutObject(object, fd)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	uploadedUrl :="oss://" + uploadCredentials.UploadBucket + "/" + object
	return &uploadedUrl, nil
}


func (clientUploader *ClientUploader)UploadBytes(dataBytes []byte)(*string, error)  {
	uploadCredentials, err := clientUploader.getUploadCredentials()

	if uploadCredentials == nil {
		return nil, err
	}

	client, err := oss.New(uploadCredentials.OssEndpoint, uploadCredentials.AccessKeyId, uploadCredentials.AccessKeySecret, oss.SecurityToken(uploadCredentials.SecurityToken))
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}
	// 获取存储空间。
	bucket, err := client.Bucket(uploadCredentials.UploadBucket)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	uuid := uuid.NewV4()
	object := uploadCredentials.UploadFolder + "/" + clientUploader.Prefix + "/" + uuid.String()
	// 上传文件流。
	err = bucket.PutObject(object, bytes.NewReader(dataBytes))
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	uploadedUrl :="oss://" + uploadCredentials.UploadBucket + "/" + object
	return &uploadedUrl, nil
}

func (clientUploader *ClientUploader)getUploadCredentials()(*uploadCredentials,  error)  {
	if clientUploader.UploadCredentials == nil || clientUploader.UploadCredentials.ExpiredTime < (time.Now().UnixNano() / 1e6) {
		 uploadCredentials, err := clientUploader.getUploadCredentialsFromServer()
         if err == nil {
			 clientUploader.UploadCredentials = uploadCredentials
		 }

		 return clientUploader.UploadCredentials, err
	}

	return clientUploader.UploadCredentials, nil
}

type uploadCredentialsResponse struct {
	Code int16 `json:"code"`
	Data uploadCredentials `json:"data"`
	Msg string `json:"msg"`
	RequestId string `json:"requestId"`

}


func (clientUploader *ClientUploader)getUploadCredentialsFromServer()(*uploadCredentials, error)  {
	uploadCredentialsRequest := green.CreateUploadCredentialsRequest();
	dataJson, _:= json.Marshal(make(map[string]string))
	uploadCredentialsRequest.SetContent(dataJson)
	var uploadCredentials *uploadCredentials
	// 发起请求并处理异常
	response, err := clientUploader.GreenClient.UploadCredentials(uploadCredentialsRequest)
	if err != nil {
		// 异常处理
		return uploadCredentials, err
	}

	uploadCredentialsResponse := uploadCredentialsResponse{}
	err = json.Unmarshal([]byte(response.GetHttpContentString()), &uploadCredentialsResponse)
	if err == nil {
		uploadCredentials = &uploadCredentialsResponse.Data
	}

	return  uploadCredentials, err
}



