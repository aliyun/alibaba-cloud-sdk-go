package dypnsapi

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

// QueryGateVerifyBillingPublic invokes the dypnsapi.QueryGateVerifyBillingPublic API synchronously
func (client *Client) QueryGateVerifyBillingPublic(request *QueryGateVerifyBillingPublicRequest) (response *QueryGateVerifyBillingPublicResponse, err error) {
	response = CreateQueryGateVerifyBillingPublicResponse()
	err = client.DoAction(request, response)
	return
}

// QueryGateVerifyBillingPublicWithChan invokes the dypnsapi.QueryGateVerifyBillingPublic API asynchronously
func (client *Client) QueryGateVerifyBillingPublicWithChan(request *QueryGateVerifyBillingPublicRequest) (<-chan *QueryGateVerifyBillingPublicResponse, <-chan error) {
	responseChan := make(chan *QueryGateVerifyBillingPublicResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryGateVerifyBillingPublic(request)
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

// QueryGateVerifyBillingPublicWithCallback invokes the dypnsapi.QueryGateVerifyBillingPublic API asynchronously
func (client *Client) QueryGateVerifyBillingPublicWithCallback(request *QueryGateVerifyBillingPublicRequest, callback func(response *QueryGateVerifyBillingPublicResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryGateVerifyBillingPublicResponse
		var err error
		defer close(result)
		response, err = client.QueryGateVerifyBillingPublic(request)
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

// QueryGateVerifyBillingPublicRequest is the request struct for api QueryGateVerifyBillingPublic
type QueryGateVerifyBillingPublicRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	AuthenticationType   requests.Integer `position:"Query" name:"AuthenticationType"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ProdCode             string           `position:"Query" name:"ProdCode"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	Month                string           `position:"Query" name:"Month"`
}

// QueryGateVerifyBillingPublicResponse is the response struct for api QueryGateVerifyBillingPublic
type QueryGateVerifyBillingPublicResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateQueryGateVerifyBillingPublicRequest creates a request to invoke QueryGateVerifyBillingPublic API
func CreateQueryGateVerifyBillingPublicRequest() (request *QueryGateVerifyBillingPublicRequest) {
	request = &QueryGateVerifyBillingPublicRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dypnsapi", "2017-05-25", "QueryGateVerifyBillingPublic", "dypnsapi", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQueryGateVerifyBillingPublicResponse creates a response to parse from QueryGateVerifyBillingPublic response
func CreateQueryGateVerifyBillingPublicResponse() (response *QueryGateVerifyBillingPublicResponse) {
	response = &QueryGateVerifyBillingPublicResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
