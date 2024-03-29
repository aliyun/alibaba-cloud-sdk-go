package cloud_siem

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

// ListAccountAccessId invokes the cloud_siem.ListAccountAccessId API synchronously
func (client *Client) ListAccountAccessId(request *ListAccountAccessIdRequest) (response *ListAccountAccessIdResponse, err error) {
	response = CreateListAccountAccessIdResponse()
	err = client.DoAction(request, response)
	return
}

// ListAccountAccessIdWithChan invokes the cloud_siem.ListAccountAccessId API asynchronously
func (client *Client) ListAccountAccessIdWithChan(request *ListAccountAccessIdRequest) (<-chan *ListAccountAccessIdResponse, <-chan error) {
	responseChan := make(chan *ListAccountAccessIdResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListAccountAccessId(request)
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

// ListAccountAccessIdWithCallback invokes the cloud_siem.ListAccountAccessId API asynchronously
func (client *Client) ListAccountAccessIdWithCallback(request *ListAccountAccessIdRequest, callback func(response *ListAccountAccessIdResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListAccountAccessIdResponse
		var err error
		defer close(result)
		response, err = client.ListAccountAccessId(request)
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

// ListAccountAccessIdRequest is the request struct for api ListAccountAccessId
type ListAccountAccessIdRequest struct {
	*requests.RpcRequest
	CloudCode string `position:"Body" name:"CloudCode"`
}

// ListAccountAccessIdResponse is the response struct for api ListAccountAccessId
type ListAccountAccessIdResponse struct {
	*responses.BaseResponse
	Success   bool       `json:"Success" xml:"Success"`
	Code      int        `json:"Code" xml:"Code"`
	Message   string     `json:"Message" xml:"Message"`
	RequestId string     `json:"RequestId" xml:"RequestId"`
	Data      []DataItem `json:"Data" xml:"Data"`
}

// CreateListAccountAccessIdRequest creates a request to invoke ListAccountAccessId API
func CreateListAccountAccessIdRequest() (request *ListAccountAccessIdRequest) {
	request = &ListAccountAccessIdRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloud-siem", "2022-06-16", "ListAccountAccessId", "cloud-siem", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListAccountAccessIdResponse creates a response to parse from ListAccountAccessId response
func CreateListAccountAccessIdResponse() (response *ListAccountAccessIdResponse) {
	response = &ListAccountAccessIdResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
