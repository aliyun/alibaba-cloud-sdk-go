package dms_enterprise

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

// DeleteScenario invokes the dms_enterprise.DeleteScenario API synchronously
func (client *Client) DeleteScenario(request *DeleteScenarioRequest) (response *DeleteScenarioResponse, err error) {
	response = CreateDeleteScenarioResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteScenarioWithChan invokes the dms_enterprise.DeleteScenario API asynchronously
func (client *Client) DeleteScenarioWithChan(request *DeleteScenarioRequest) (<-chan *DeleteScenarioResponse, <-chan error) {
	responseChan := make(chan *DeleteScenarioResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteScenario(request)
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

// DeleteScenarioWithCallback invokes the dms_enterprise.DeleteScenario API asynchronously
func (client *Client) DeleteScenarioWithCallback(request *DeleteScenarioRequest, callback func(response *DeleteScenarioResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteScenarioResponse
		var err error
		defer close(result)
		response, err = client.DeleteScenario(request)
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

// DeleteScenarioRequest is the request struct for api DeleteScenario
type DeleteScenarioRequest struct {
	*requests.RpcRequest
	Tid        requests.Integer `position:"Query" name:"Tid"`
	ScenarioId requests.Integer `position:"Query" name:"ScenarioId"`
}

// DeleteScenarioResponse is the response struct for api DeleteScenario
type DeleteScenarioResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	ErrorCode    string `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	Success      bool   `json:"Success" xml:"Success"`
}

// CreateDeleteScenarioRequest creates a request to invoke DeleteScenario API
func CreateDeleteScenarioRequest() (request *DeleteScenarioRequest) {
	request = &DeleteScenarioRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dms-enterprise", "2018-11-01", "DeleteScenario", "dms-enterprise", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteScenarioResponse creates a response to parse from DeleteScenario response
func CreateDeleteScenarioResponse() (response *DeleteScenarioResponse) {
	response = &DeleteScenarioResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
