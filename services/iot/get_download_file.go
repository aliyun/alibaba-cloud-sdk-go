package iot

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

// GetDownloadFile invokes the iot.GetDownloadFile API synchronously
func (client *Client) GetDownloadFile(request *GetDownloadFileRequest) (response *GetDownloadFileResponse, err error) {
	response = CreateGetDownloadFileResponse()
	err = client.DoAction(request, response)
	return
}

// GetDownloadFileWithChan invokes the iot.GetDownloadFile API asynchronously
func (client *Client) GetDownloadFileWithChan(request *GetDownloadFileRequest) (<-chan *GetDownloadFileResponse, <-chan error) {
	responseChan := make(chan *GetDownloadFileResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetDownloadFile(request)
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

// GetDownloadFileWithCallback invokes the iot.GetDownloadFile API asynchronously
func (client *Client) GetDownloadFileWithCallback(request *GetDownloadFileRequest, callback func(response *GetDownloadFileResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetDownloadFileResponse
		var err error
		defer close(result)
		response, err = client.GetDownloadFile(request)
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

// GetDownloadFileRequest is the request struct for api GetDownloadFile
type GetDownloadFileRequest struct {
	*requests.RpcRequest
	IotInstanceId string `position:"Body" name:"IotInstanceId"`
	Context       string `position:"Body" name:"Context"`
	LongJobId     string `position:"Query" name:"LongJobId"`
	ApiProduct    string `position:"Body" name:"ApiProduct"`
	ApiRevision   string `position:"Body" name:"ApiRevision"`
}

// GetDownloadFileResponse is the response struct for api GetDownloadFile
type GetDownloadFileResponse struct {
	*responses.BaseResponse
	RequestId    string                `json:"RequestId" xml:"RequestId"`
	Success      bool                  `json:"Success" xml:"Success"`
	Code         string                `json:"Code" xml:"Code"`
	ErrorMessage string                `json:"ErrorMessage" xml:"ErrorMessage"`
	Data         DataInGetDownloadFile `json:"Data" xml:"Data"`
}

// CreateGetDownloadFileRequest creates a request to invoke GetDownloadFile API
func CreateGetDownloadFileRequest() (request *GetDownloadFileRequest) {
	request = &GetDownloadFileRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "GetDownloadFile", "iot", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetDownloadFileResponse creates a response to parse from GetDownloadFile response
func CreateGetDownloadFileResponse() (response *GetDownloadFileResponse) {
	response = &GetDownloadFileResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
