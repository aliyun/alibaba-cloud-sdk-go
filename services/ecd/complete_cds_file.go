package ecd

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

// CompleteCdsFile invokes the ecd.CompleteCdsFile API synchronously
func (client *Client) CompleteCdsFile(request *CompleteCdsFileRequest) (response *CompleteCdsFileResponse, err error) {
	response = CreateCompleteCdsFileResponse()
	err = client.DoAction(request, response)
	return
}

// CompleteCdsFileWithChan invokes the ecd.CompleteCdsFile API asynchronously
func (client *Client) CompleteCdsFileWithChan(request *CompleteCdsFileRequest) (<-chan *CompleteCdsFileResponse, <-chan error) {
	responseChan := make(chan *CompleteCdsFileResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CompleteCdsFile(request)
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

// CompleteCdsFileWithCallback invokes the ecd.CompleteCdsFile API asynchronously
func (client *Client) CompleteCdsFileWithCallback(request *CompleteCdsFileRequest, callback func(response *CompleteCdsFileResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CompleteCdsFileResponse
		var err error
		defer close(result)
		response, err = client.CompleteCdsFile(request)
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

// CompleteCdsFileRequest is the request struct for api CompleteCdsFile
type CompleteCdsFileRequest struct {
	*requests.RpcRequest
	UploadId  string `position:"Query" name:"UploadId"`
	GroupId   string `position:"Query" name:"GroupId"`
	CdsId     string `position:"Query" name:"CdsId"`
	EndUserId string `position:"Query" name:"EndUserId"`
	FileId    string `position:"Query" name:"FileId"`
}

// CompleteCdsFileResponse is the response struct for api CompleteCdsFile
type CompleteCdsFileResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateCompleteCdsFileRequest creates a request to invoke CompleteCdsFile API
func CreateCompleteCdsFileRequest() (request *CompleteCdsFileRequest) {
	request = &CompleteCdsFileRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ecd", "2020-09-30", "CompleteCdsFile", "gwsecd", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCompleteCdsFileResponse creates a response to parse from CompleteCdsFile response
func CreateCompleteCdsFileResponse() (response *CompleteCdsFileResponse) {
	response = &CompleteCdsFileResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
