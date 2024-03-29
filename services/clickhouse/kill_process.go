package clickhouse

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

// KillProcess invokes the clickhouse.KillProcess API synchronously
func (client *Client) KillProcess(request *KillProcessRequest) (response *KillProcessResponse, err error) {
	response = CreateKillProcessResponse()
	err = client.DoAction(request, response)
	return
}

// KillProcessWithChan invokes the clickhouse.KillProcess API asynchronously
func (client *Client) KillProcessWithChan(request *KillProcessRequest) (<-chan *KillProcessResponse, <-chan error) {
	responseChan := make(chan *KillProcessResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.KillProcess(request)
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

// KillProcessWithCallback invokes the clickhouse.KillProcess API asynchronously
func (client *Client) KillProcessWithCallback(request *KillProcessRequest, callback func(response *KillProcessResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *KillProcessResponse
		var err error
		defer close(result)
		response, err = client.KillProcess(request)
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

// KillProcessRequest is the request struct for api KillProcess
type KillProcessRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	InitialQueryId       string           `position:"Query" name:"InitialQueryId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	DBClusterId          string           `position:"Query" name:"DBClusterId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// KillProcessResponse is the response struct for api KillProcess
type KillProcessResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateKillProcessRequest creates a request to invoke KillProcess API
func CreateKillProcessRequest() (request *KillProcessRequest) {
	request = &KillProcessRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("clickhouse", "2019-11-11", "KillProcess", "service", "openAPI")
	request.Method = requests.POST
	return
}

// CreateKillProcessResponse creates a response to parse from KillProcess response
func CreateKillProcessResponse() (response *KillProcessResponse) {
	response = &KillProcessResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
