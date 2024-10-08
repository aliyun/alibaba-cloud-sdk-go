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

// DeployFile invokes the dataworks_public.DeployFile API synchronously
func (client *Client) DeployFile(request *DeployFileRequest) (response *DeployFileResponse, err error) {
	response = CreateDeployFileResponse()
	err = client.DoAction(request, response)
	return
}

// DeployFileWithChan invokes the dataworks_public.DeployFile API asynchronously
func (client *Client) DeployFileWithChan(request *DeployFileRequest) (<-chan *DeployFileResponse, <-chan error) {
	responseChan := make(chan *DeployFileResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeployFile(request)
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

// DeployFileWithCallback invokes the dataworks_public.DeployFile API asynchronously
func (client *Client) DeployFileWithCallback(request *DeployFileRequest, callback func(response *DeployFileResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeployFileResponse
		var err error
		defer close(result)
		response, err = client.DeployFile(request)
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

// DeployFileRequest is the request struct for api DeployFile
type DeployFileRequest struct {
	*requests.RpcRequest
	ProjectIdentifier string           `position:"Body" name:"ProjectIdentifier"`
	Comment           string           `position:"Body" name:"Comment"`
	ProjectId         requests.Integer `position:"Body" name:"ProjectId"`
	NodeId            requests.Integer `position:"Body" name:"NodeId"`
	FileId            requests.Integer `position:"Body" name:"FileId"`
}

// DeployFileResponse is the response struct for api DeployFile
type DeployFileResponse struct {
	*responses.BaseResponse
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Data           int64  `json:"Data" xml:"Data"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	ErrorMessage   string `json:"ErrorMessage" xml:"ErrorMessage"`
	Success        bool   `json:"Success" xml:"Success"`
	ErrorCode      string `json:"ErrorCode" xml:"ErrorCode"`
}

// CreateDeployFileRequest creates a request to invoke DeployFile API
func CreateDeployFileRequest() (request *DeployFileRequest) {
	request = &DeployFileRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dataworks-public", "2020-05-18", "DeployFile", "", "")
	request.Method = requests.POST
	return
}

// CreateDeployFileResponse creates a response to parse from DeployFile response
func CreateDeployFileResponse() (response *DeployFileResponse) {
	response = &DeployFileResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
