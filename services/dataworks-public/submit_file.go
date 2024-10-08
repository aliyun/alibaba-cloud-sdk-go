package dataworks_public

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

// SubmitFile invokes the dataworks_public.SubmitFile API synchronously
func (client *Client) SubmitFile(request *SubmitFileRequest) (response *SubmitFileResponse, err error) {
	response = CreateSubmitFileResponse()
	err = client.DoAction(request, response)
	return
}

// SubmitFileWithChan invokes the dataworks_public.SubmitFile API asynchronously
func (client *Client) SubmitFileWithChan(request *SubmitFileRequest) (<-chan *SubmitFileResponse, <-chan error) {
	responseChan := make(chan *SubmitFileResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SubmitFile(request)
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

// SubmitFileWithCallback invokes the dataworks_public.SubmitFile API asynchronously
func (client *Client) SubmitFileWithCallback(request *SubmitFileRequest, callback func(response *SubmitFileResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SubmitFileResponse
		var err error
		defer close(result)
		response, err = client.SubmitFile(request)
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

// SubmitFileRequest is the request struct for api SubmitFile
type SubmitFileRequest struct {
	*requests.RpcRequest
	ProjectIdentifier           string           `position:"Body" name:"ProjectIdentifier"`
	SkipAllDeployFileExtensions requests.Boolean `position:"Body" name:"SkipAllDeployFileExtensions"`
	Comment                     string           `position:"Body" name:"Comment"`
	ProjectId                   requests.Integer `position:"Body" name:"ProjectId"`
	FileId                      requests.Integer `position:"Body" name:"FileId"`
}

// SubmitFileResponse is the response struct for api SubmitFile
type SubmitFileResponse struct {
	*responses.BaseResponse
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Data           int64  `json:"Data" xml:"Data"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	ErrorMessage   string `json:"ErrorMessage" xml:"ErrorMessage"`
	Success        bool   `json:"Success" xml:"Success"`
	ErrorCode      string `json:"ErrorCode" xml:"ErrorCode"`
}

// CreateSubmitFileRequest creates a request to invoke SubmitFile API
func CreateSubmitFileRequest() (request *SubmitFileRequest) {
	request = &SubmitFileRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dataworks-public", "2020-05-18", "SubmitFile", "", "")
	request.Method = requests.POST
	return
}

// CreateSubmitFileResponse creates a response to parse from SubmitFile response
func CreateSubmitFileResponse() (response *SubmitFileResponse) {
	response = &SubmitFileResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
