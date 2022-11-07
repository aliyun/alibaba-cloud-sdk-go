package arms

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

// DeleteEventBridgeIntegration invokes the arms.DeleteEventBridgeIntegration API synchronously
func (client *Client) DeleteEventBridgeIntegration(request *DeleteEventBridgeIntegrationRequest) (response *DeleteEventBridgeIntegrationResponse, err error) {
	response = CreateDeleteEventBridgeIntegrationResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteEventBridgeIntegrationWithChan invokes the arms.DeleteEventBridgeIntegration API asynchronously
func (client *Client) DeleteEventBridgeIntegrationWithChan(request *DeleteEventBridgeIntegrationRequest) (<-chan *DeleteEventBridgeIntegrationResponse, <-chan error) {
	responseChan := make(chan *DeleteEventBridgeIntegrationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteEventBridgeIntegration(request)
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

// DeleteEventBridgeIntegrationWithCallback invokes the arms.DeleteEventBridgeIntegration API asynchronously
func (client *Client) DeleteEventBridgeIntegrationWithCallback(request *DeleteEventBridgeIntegrationRequest, callback func(response *DeleteEventBridgeIntegrationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteEventBridgeIntegrationResponse
		var err error
		defer close(result)
		response, err = client.DeleteEventBridgeIntegration(request)
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

// DeleteEventBridgeIntegrationRequest is the request struct for api DeleteEventBridgeIntegration
type DeleteEventBridgeIntegrationRequest struct {
	*requests.RpcRequest
	Id requests.Integer `position:"Body" name:"Id"`
}

// DeleteEventBridgeIntegrationResponse is the response struct for api DeleteEventBridgeIntegration
type DeleteEventBridgeIntegrationResponse struct {
	*responses.BaseResponse
	RequestId                             string `json:"RequestId" xml:"RequestId"`
	DeleteEventBridgeIntegrationIsSuccess bool   `json:"IsSuccess" xml:"IsSuccess"`
}

// CreateDeleteEventBridgeIntegrationRequest creates a request to invoke DeleteEventBridgeIntegration API
func CreateDeleteEventBridgeIntegrationRequest() (request *DeleteEventBridgeIntegrationRequest) {
	request = &DeleteEventBridgeIntegrationRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ARMS", "2019-08-08", "DeleteEventBridgeIntegration", "arms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteEventBridgeIntegrationResponse creates a response to parse from DeleteEventBridgeIntegration response
func CreateDeleteEventBridgeIntegrationResponse() (response *DeleteEventBridgeIntegrationResponse) {
	response = &DeleteEventBridgeIntegrationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
