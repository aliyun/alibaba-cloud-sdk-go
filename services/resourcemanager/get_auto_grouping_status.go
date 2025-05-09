package resourcemanager

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

// GetAutoGroupingStatus invokes the resourcemanager.GetAutoGroupingStatus API synchronously
func (client *Client) GetAutoGroupingStatus(request *GetAutoGroupingStatusRequest) (response *GetAutoGroupingStatusResponse, err error) {
	response = CreateGetAutoGroupingStatusResponse()
	err = client.DoAction(request, response)
	return
}

// GetAutoGroupingStatusWithChan invokes the resourcemanager.GetAutoGroupingStatus API asynchronously
func (client *Client) GetAutoGroupingStatusWithChan(request *GetAutoGroupingStatusRequest) (<-chan *GetAutoGroupingStatusResponse, <-chan error) {
	responseChan := make(chan *GetAutoGroupingStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetAutoGroupingStatus(request)
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

// GetAutoGroupingStatusWithCallback invokes the resourcemanager.GetAutoGroupingStatus API asynchronously
func (client *Client) GetAutoGroupingStatusWithCallback(request *GetAutoGroupingStatusRequest, callback func(response *GetAutoGroupingStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetAutoGroupingStatusResponse
		var err error
		defer close(result)
		response, err = client.GetAutoGroupingStatus(request)
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

// GetAutoGroupingStatusRequest is the request struct for api GetAutoGroupingStatus
type GetAutoGroupingStatusRequest struct {
	*requests.RpcRequest
}

// GetAutoGroupingStatusResponse is the response struct for api GetAutoGroupingStatus
type GetAutoGroupingStatusResponse struct {
	*responses.BaseResponse
	EnableExistedResourcesTransfer bool   `json:"EnableExistedResourcesTransfer" xml:"EnableExistedResourcesTransfer"`
	EnableStatus                   string `json:"EnableStatus" xml:"EnableStatus"`
	RequestId                      string `json:"RequestId" xml:"RequestId"`
}

// CreateGetAutoGroupingStatusRequest creates a request to invoke GetAutoGroupingStatus API
func CreateGetAutoGroupingStatusRequest() (request *GetAutoGroupingStatusRequest) {
	request = &GetAutoGroupingStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ResourceManager", "2020-03-31", "GetAutoGroupingStatus", "resourcemanager", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetAutoGroupingStatusResponse creates a response to parse from GetAutoGroupingStatus response
func CreateGetAutoGroupingStatusResponse() (response *GetAutoGroupingStatusResponse) {
	response = &GetAutoGroupingStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
