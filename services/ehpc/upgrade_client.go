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

// UpgradeClient invokes the ehpc.UpgradeClient API synchronously
func (client *Client) UpgradeClient(request *UpgradeClientRequest) (response *UpgradeClientResponse, err error) {
	response = CreateUpgradeClientResponse()
	err = client.DoAction(request, response)
	return
}

// UpgradeClientWithChan invokes the ehpc.UpgradeClient API asynchronously
func (client *Client) UpgradeClientWithChan(request *UpgradeClientRequest) (<-chan *UpgradeClientResponse, <-chan error) {
	responseChan := make(chan *UpgradeClientResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpgradeClient(request)
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

// UpgradeClientWithCallback invokes the ehpc.UpgradeClient API asynchronously
func (client *Client) UpgradeClientWithCallback(request *UpgradeClientRequest, callback func(response *UpgradeClientResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpgradeClientResponse
		var err error
		defer close(result)
		response, err = client.UpgradeClient(request)
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

// UpgradeClientRequest is the request struct for api UpgradeClient
type UpgradeClientRequest struct {
	*requests.RpcRequest
	ClientVersion string `position:"Query" name:"ClientVersion"`
	ClusterId     string `position:"Query" name:"ClusterId"`
}

// UpgradeClientResponse is the response struct for api UpgradeClient
type UpgradeClientResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUpgradeClientRequest creates a request to invoke UpgradeClient API
func CreateUpgradeClientRequest() (request *UpgradeClientRequest) {
	request = &UpgradeClientRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("EHPC", "2018-04-12", "UpgradeClient", "ehs", "openAPI")
	request.Method = requests.GET
	return
}

// CreateUpgradeClientResponse creates a response to parse from UpgradeClient response
func CreateUpgradeClientResponse() (response *UpgradeClientResponse) {
	response = &UpgradeClientResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
