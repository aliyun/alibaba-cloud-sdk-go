package resourcesharing

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

// EnableSharingWithResourceDirectory invokes the resourcesharing.EnableSharingWithResourceDirectory API synchronously
func (client *Client) EnableSharingWithResourceDirectory(request *EnableSharingWithResourceDirectoryRequest) (response *EnableSharingWithResourceDirectoryResponse, err error) {
	response = CreateEnableSharingWithResourceDirectoryResponse()
	err = client.DoAction(request, response)
	return
}

// EnableSharingWithResourceDirectoryWithChan invokes the resourcesharing.EnableSharingWithResourceDirectory API asynchronously
func (client *Client) EnableSharingWithResourceDirectoryWithChan(request *EnableSharingWithResourceDirectoryRequest) (<-chan *EnableSharingWithResourceDirectoryResponse, <-chan error) {
	responseChan := make(chan *EnableSharingWithResourceDirectoryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.EnableSharingWithResourceDirectory(request)
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

// EnableSharingWithResourceDirectoryWithCallback invokes the resourcesharing.EnableSharingWithResourceDirectory API asynchronously
func (client *Client) EnableSharingWithResourceDirectoryWithCallback(request *EnableSharingWithResourceDirectoryRequest, callback func(response *EnableSharingWithResourceDirectoryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *EnableSharingWithResourceDirectoryResponse
		var err error
		defer close(result)
		response, err = client.EnableSharingWithResourceDirectory(request)
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

// EnableSharingWithResourceDirectoryRequest is the request struct for api EnableSharingWithResourceDirectory
type EnableSharingWithResourceDirectoryRequest struct {
	*requests.RpcRequest
}

// EnableSharingWithResourceDirectoryResponse is the response struct for api EnableSharingWithResourceDirectory
type EnableSharingWithResourceDirectoryResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateEnableSharingWithResourceDirectoryRequest creates a request to invoke EnableSharingWithResourceDirectory API
func CreateEnableSharingWithResourceDirectoryRequest() (request *EnableSharingWithResourceDirectoryRequest) {
	request = &EnableSharingWithResourceDirectoryRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ResourceSharing", "2020-01-10", "EnableSharingWithResourceDirectory", "ressharing", "openAPI")
	request.Method = requests.POST
	return
}

// CreateEnableSharingWithResourceDirectoryResponse creates a response to parse from EnableSharingWithResourceDirectory response
func CreateEnableSharingWithResourceDirectoryResponse() (response *EnableSharingWithResourceDirectoryResponse) {
	response = &EnableSharingWithResourceDirectoryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
