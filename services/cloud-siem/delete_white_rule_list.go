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

// DeleteWhiteRuleList invokes the cloud_siem.DeleteWhiteRuleList API synchronously
func (client *Client) DeleteWhiteRuleList(request *DeleteWhiteRuleListRequest) (response *DeleteWhiteRuleListResponse, err error) {
	response = CreateDeleteWhiteRuleListResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteWhiteRuleListWithChan invokes the cloud_siem.DeleteWhiteRuleList API asynchronously
func (client *Client) DeleteWhiteRuleListWithChan(request *DeleteWhiteRuleListRequest) (<-chan *DeleteWhiteRuleListResponse, <-chan error) {
	responseChan := make(chan *DeleteWhiteRuleListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteWhiteRuleList(request)
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

// DeleteWhiteRuleListWithCallback invokes the cloud_siem.DeleteWhiteRuleList API asynchronously
func (client *Client) DeleteWhiteRuleListWithCallback(request *DeleteWhiteRuleListRequest, callback func(response *DeleteWhiteRuleListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteWhiteRuleListResponse
		var err error
		defer close(result)
		response, err = client.DeleteWhiteRuleList(request)
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

// DeleteWhiteRuleListRequest is the request struct for api DeleteWhiteRuleList
type DeleteWhiteRuleListRequest struct {
	*requests.RpcRequest
	Id requests.Integer `position:"Body" name:"Id"`
}

// DeleteWhiteRuleListResponse is the response struct for api DeleteWhiteRuleList
type DeleteWhiteRuleListResponse struct {
	*responses.BaseResponse
}

// CreateDeleteWhiteRuleListRequest creates a request to invoke DeleteWhiteRuleList API
func CreateDeleteWhiteRuleListRequest() (request *DeleteWhiteRuleListRequest) {
	request = &DeleteWhiteRuleListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloud-siem", "2022-06-16", "DeleteWhiteRuleList", "cloud-siem", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteWhiteRuleListResponse creates a response to parse from DeleteWhiteRuleList response
func CreateDeleteWhiteRuleListResponse() (response *DeleteWhiteRuleListResponse) {
	response = &DeleteWhiteRuleListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
