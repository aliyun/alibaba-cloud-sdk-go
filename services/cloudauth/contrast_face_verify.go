package cloudauth

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// ContrastFaceVerify invokes the cloudauth.ContrastFaceVerify API synchronously
func (client *Client) ContrastFaceVerify(request *ContrastFaceVerifyRequest) (response *ContrastFaceVerifyResponse, err error) {
	response = CreateContrastFaceVerifyResponse()
	err = client.DoAction(request, response)
	return
}

// ContrastFaceVerifyWithChan invokes the cloudauth.ContrastFaceVerify API asynchronously
func (client *Client) ContrastFaceVerifyWithChan(request *ContrastFaceVerifyRequest) (<-chan *ContrastFaceVerifyResponse, <-chan error) {
	responseChan := make(chan *ContrastFaceVerifyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ContrastFaceVerify(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// ContrastFaceVerifyWithCallback invokes the cloudauth.ContrastFaceVerify API asynchronously
func (client *Client) ContrastFaceVerifyWithCallback(request *ContrastFaceVerifyRequest, callback func(response *ContrastFaceVerifyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ContrastFaceVerifyResponse
		var err error
		defer close(result)
		response, err = client.ContrastFaceVerify(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// ContrastFaceVerifyRequest is the request struct for api ContrastFaceVerify
type ContrastFaceVerifyRequest struct {
	*requests.RpcRequest
	ProductCode            string           `position:"Body" name:"ProductCode"`
	FaceContrastPicture    string           `position:"Body" name:"FaceContrastPicture"`
	DeviceToken            string           `position:"Body" name:"DeviceToken"`
	UserId                 string           `position:"Body" name:"UserId"`
	CertifyId              string           `position:"Body" name:"CertifyId"`
	EncryptType            string           `position:"Body" name:"EncryptType"`
	CertNo                 string           `position:"Body" name:"CertNo"`
	OuterOrderNo           string           `position:"Body" name:"OuterOrderNo"`
	CertType               string           `position:"Body" name:"CertType"`
	FaceContrastPictureUrl string           `position:"Body" name:"FaceContrastPictureUrl"`
	Model                  string           `position:"Query" name:"Model"`
	OssObjectName          string           `position:"Body" name:"OssObjectName"`
	CertName               string           `position:"Body" name:"CertName"`
	Ip                     string           `position:"Body" name:"Ip"`
	Mobile                 string           `position:"Body" name:"Mobile"`
	FaceContrastFile       string           `position:"Body" name:"FaceContrastFile"`
	SceneId                requests.Integer `position:"Body" name:"SceneId"`
	OssBucketName          string           `position:"Body" name:"OssBucketName"`
	Crop                   string           `position:"Body" name:"Crop"`
}

// ContrastFaceVerifyResponse is the response struct for api ContrastFaceVerify
type ContrastFaceVerifyResponse struct {
	*responses.BaseResponse
	Code         string       `json:"Code" xml:"Code"`
	Message      string       `json:"Message" xml:"Message"`
	RequestId    string       `json:"RequestId" xml:"RequestId"`
	ResultObject ResultObject `json:"ResultObject" xml:"ResultObject"`
}

// CreateContrastFaceVerifyRequest creates a request to invoke ContrastFaceVerify API
func CreateContrastFaceVerifyRequest() (request *ContrastFaceVerifyRequest) {
	request = &ContrastFaceVerifyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cloudauth", "2019-03-07", "ContrastFaceVerify", "cloudauth", "openAPI")
	request.Method = requests.POST
	return
}

// CreateContrastFaceVerifyResponse creates a response to parse from ContrastFaceVerify response
func CreateContrastFaceVerifyResponse() (response *ContrastFaceVerifyResponse) {
	response = &ContrastFaceVerifyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
