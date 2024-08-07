package elasticsearch

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

// RestartInstance invokes the elasticsearch.RestartInstance API synchronously
func (client *Client) RestartInstance(request *RestartInstanceRequest) (response *RestartInstanceResponse, err error) {
	response = CreateRestartInstanceResponse()
	err = client.DoAction(request, response)
	return
}

// RestartInstanceWithChan invokes the elasticsearch.RestartInstance API asynchronously
func (client *Client) RestartInstanceWithChan(request *RestartInstanceRequest) (<-chan *RestartInstanceResponse, <-chan error) {
	responseChan := make(chan *RestartInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RestartInstance(request)
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

// RestartInstanceWithCallback invokes the elasticsearch.RestartInstance API asynchronously
func (client *Client) RestartInstanceWithCallback(request *RestartInstanceRequest, callback func(response *RestartInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RestartInstanceResponse
		var err error
		defer close(result)
		response, err = client.RestartInstance(request)
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

// RestartInstanceRequest is the request struct for api RestartInstance
type RestartInstanceRequest struct {
	*requests.RoaRequest
	InstanceId  string           `position:"Path" name:"InstanceId"`
	ClientToken string           `position:"Query" name:"clientToken"`
	Force       requests.Boolean `position:"Query" name:"force"`
	Body        string           `position:"Body" name:"body"`
}

// RestartInstanceResponse is the response struct for api RestartInstance
type RestartInstanceResponse struct {
	*responses.BaseResponse
	RequestId string                  `json:"RequestId" xml:"RequestId"`
	Result    ResultInRestartInstance `json:"Result" xml:"Result"`
}

// CreateRestartInstanceRequest creates a request to invoke RestartInstance API
func CreateRestartInstanceRequest() (request *RestartInstanceRequest) {
	request = &RestartInstanceRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("elasticsearch", "2017-06-13", "RestartInstance", "/openapi/instances/[InstanceId]/actions/restart", "elasticsearch", "openAPI")
	request.Method = requests.POST
	return
}

// CreateRestartInstanceResponse creates a response to parse from RestartInstance response
func CreateRestartInstanceResponse() (response *RestartInstanceResponse) {
	response = &RestartInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
