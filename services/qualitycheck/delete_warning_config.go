package qualitycheck

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

// DeleteWarningConfig invokes the qualitycheck.DeleteWarningConfig API synchronously
func (client *Client) DeleteWarningConfig(request *DeleteWarningConfigRequest) (response *DeleteWarningConfigResponse, err error) {
	response = CreateDeleteWarningConfigResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteWarningConfigWithChan invokes the qualitycheck.DeleteWarningConfig API asynchronously
func (client *Client) DeleteWarningConfigWithChan(request *DeleteWarningConfigRequest) (<-chan *DeleteWarningConfigResponse, <-chan error) {
	responseChan := make(chan *DeleteWarningConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteWarningConfig(request)
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

// DeleteWarningConfigWithCallback invokes the qualitycheck.DeleteWarningConfig API asynchronously
func (client *Client) DeleteWarningConfigWithCallback(request *DeleteWarningConfigRequest, callback func(response *DeleteWarningConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteWarningConfigResponse
		var err error
		defer close(result)
		response, err = client.DeleteWarningConfig(request)
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

// DeleteWarningConfigRequest is the request struct for api DeleteWarningConfig
type DeleteWarningConfigRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	JsonStr         string           `position:"Query" name:"JsonStr"`
	BaseMeAgentId   requests.Integer `position:"Query" name:"BaseMeAgentId"`
}

// DeleteWarningConfigResponse is the response struct for api DeleteWarningConfig
type DeleteWarningConfigResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateDeleteWarningConfigRequest creates a request to invoke DeleteWarningConfig API
func CreateDeleteWarningConfigRequest() (request *DeleteWarningConfigRequest) {
	request = &DeleteWarningConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Qualitycheck", "2019-01-15", "DeleteWarningConfig", "", "")
	request.Method = requests.POST
	return
}

// CreateDeleteWarningConfigResponse creates a response to parse from DeleteWarningConfig response
func CreateDeleteWarningConfigResponse() (response *DeleteWarningConfigResponse) {
	response = &DeleteWarningConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
