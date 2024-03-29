package ehpc

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

// ListContainerApps invokes the ehpc.ListContainerApps API synchronously
func (client *Client) ListContainerApps(request *ListContainerAppsRequest) (response *ListContainerAppsResponse, err error) {
	response = CreateListContainerAppsResponse()
	err = client.DoAction(request, response)
	return
}

// ListContainerAppsWithChan invokes the ehpc.ListContainerApps API asynchronously
func (client *Client) ListContainerAppsWithChan(request *ListContainerAppsRequest) (<-chan *ListContainerAppsResponse, <-chan error) {
	responseChan := make(chan *ListContainerAppsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListContainerApps(request)
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

// ListContainerAppsWithCallback invokes the ehpc.ListContainerApps API asynchronously
func (client *Client) ListContainerAppsWithCallback(request *ListContainerAppsRequest, callback func(response *ListContainerAppsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListContainerAppsResponse
		var err error
		defer close(result)
		response, err = client.ListContainerApps(request)
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

// ListContainerAppsRequest is the request struct for api ListContainerApps
type ListContainerAppsRequest struct {
	*requests.RpcRequest
	PageNumber requests.Integer `position:"Query" name:"PageNumber"`
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
}

// ListContainerAppsResponse is the response struct for api ListContainerApps
type ListContainerAppsResponse struct {
	*responses.BaseResponse
	PageSize      int           `json:"PageSize" xml:"PageSize"`
	RequestId     string        `json:"RequestId" xml:"RequestId"`
	PageNumber    int           `json:"PageNumber" xml:"PageNumber"`
	TotalCount    int           `json:"TotalCount" xml:"TotalCount"`
	ContainerApps ContainerApps `json:"ContainerApps" xml:"ContainerApps"`
}

// CreateListContainerAppsRequest creates a request to invoke ListContainerApps API
func CreateListContainerAppsRequest() (request *ListContainerAppsRequest) {
	request = &ListContainerAppsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("EHPC", "2018-04-12", "ListContainerApps", "ehs", "openAPI")
	request.Method = requests.GET
	return
}

// CreateListContainerAppsResponse creates a response to parse from ListContainerApps response
func CreateListContainerAppsResponse() (response *ListContainerAppsResponse) {
	response = &ListContainerAppsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
