package cbn

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

// DeactiveFlowLog invokes the cbn.DeactiveFlowLog API synchronously
func (client *Client) DeactiveFlowLog(request *DeactiveFlowLogRequest) (response *DeactiveFlowLogResponse, err error) {
	response = CreateDeactiveFlowLogResponse()
	err = client.DoAction(request, response)
	return
}

// DeactiveFlowLogWithChan invokes the cbn.DeactiveFlowLog API asynchronously
func (client *Client) DeactiveFlowLogWithChan(request *DeactiveFlowLogRequest) (<-chan *DeactiveFlowLogResponse, <-chan error) {
	responseChan := make(chan *DeactiveFlowLogResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeactiveFlowLog(request)
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

// DeactiveFlowLogWithCallback invokes the cbn.DeactiveFlowLog API asynchronously
func (client *Client) DeactiveFlowLogWithCallback(request *DeactiveFlowLogRequest, callback func(response *DeactiveFlowLogResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeactiveFlowLogResponse
		var err error
		defer close(result)
		response, err = client.DeactiveFlowLog(request)
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

// DeactiveFlowLogRequest is the request struct for api DeactiveFlowLog
type DeactiveFlowLogRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	CenId                string           `position:"Query" name:"CenId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	Version              string           `position:"Query" name:"Version"`
	FlowLogId            string           `position:"Query" name:"FlowLogId"`
}

// DeactiveFlowLogResponse is the response struct for api DeactiveFlowLog
type DeactiveFlowLogResponse struct {
	*responses.BaseResponse
	Success   string `json:"Success" xml:"Success"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeactiveFlowLogRequest creates a request to invoke DeactiveFlowLog API
func CreateDeactiveFlowLogRequest() (request *DeactiveFlowLogRequest) {
	request = &DeactiveFlowLogRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cbn", "2017-09-12", "DeactiveFlowLog", "cbn", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeactiveFlowLogResponse creates a response to parse from DeactiveFlowLog response
func CreateDeactiveFlowLogResponse() (response *DeactiveFlowLogResponse) {
	response = &DeactiveFlowLogResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
