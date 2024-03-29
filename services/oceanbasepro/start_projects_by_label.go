package oceanbasepro

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

// StartProjectsByLabel invokes the oceanbasepro.StartProjectsByLabel API synchronously
func (client *Client) StartProjectsByLabel(request *StartProjectsByLabelRequest) (response *StartProjectsByLabelResponse, err error) {
	response = CreateStartProjectsByLabelResponse()
	err = client.DoAction(request, response)
	return
}

// StartProjectsByLabelWithChan invokes the oceanbasepro.StartProjectsByLabel API asynchronously
func (client *Client) StartProjectsByLabelWithChan(request *StartProjectsByLabelRequest) (<-chan *StartProjectsByLabelResponse, <-chan error) {
	responseChan := make(chan *StartProjectsByLabelResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.StartProjectsByLabel(request)
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

// StartProjectsByLabelWithCallback invokes the oceanbasepro.StartProjectsByLabel API asynchronously
func (client *Client) StartProjectsByLabelWithCallback(request *StartProjectsByLabelRequest, callback func(response *StartProjectsByLabelResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *StartProjectsByLabelResponse
		var err error
		defer close(result)
		response, err = client.StartProjectsByLabel(request)
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

// StartProjectsByLabelRequest is the request struct for api StartProjectsByLabel
type StartProjectsByLabelRequest struct {
	*requests.RpcRequest
	Id string `position:"Body" name:"Id"`
}

// StartProjectsByLabelResponse is the response struct for api StartProjectsByLabel
type StartProjectsByLabelResponse struct {
	*responses.BaseResponse
}

// CreateStartProjectsByLabelRequest creates a request to invoke StartProjectsByLabel API
func CreateStartProjectsByLabelRequest() (request *StartProjectsByLabelRequest) {
	request = &StartProjectsByLabelRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("OceanBasePro", "2019-09-01", "StartProjectsByLabel", "oceanbase", "openAPI")
	request.Method = requests.POST
	return
}

// CreateStartProjectsByLabelResponse creates a response to parse from StartProjectsByLabel response
func CreateStartProjectsByLabelResponse() (response *StartProjectsByLabelResponse) {
	response = &StartProjectsByLabelResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
