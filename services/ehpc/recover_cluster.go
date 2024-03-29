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

// RecoverCluster invokes the ehpc.RecoverCluster API synchronously
func (client *Client) RecoverCluster(request *RecoverClusterRequest) (response *RecoverClusterResponse, err error) {
	response = CreateRecoverClusterResponse()
	err = client.DoAction(request, response)
	return
}

// RecoverClusterWithChan invokes the ehpc.RecoverCluster API asynchronously
func (client *Client) RecoverClusterWithChan(request *RecoverClusterRequest) (<-chan *RecoverClusterResponse, <-chan error) {
	responseChan := make(chan *RecoverClusterResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RecoverCluster(request)
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

// RecoverClusterWithCallback invokes the ehpc.RecoverCluster API asynchronously
func (client *Client) RecoverClusterWithCallback(request *RecoverClusterRequest, callback func(response *RecoverClusterResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RecoverClusterResponse
		var err error
		defer close(result)
		response, err = client.RecoverCluster(request)
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

// RecoverClusterRequest is the request struct for api RecoverCluster
type RecoverClusterRequest struct {
	*requests.RpcRequest
	ImageId         string `position:"Query" name:"ImageId"`
	OsTag           string `position:"Query" name:"OsTag"`
	ClientVersion   string `position:"Query" name:"ClientVersion"`
	AccountType     string `position:"Query" name:"AccountType"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	ImageOwnerAlias string `position:"Query" name:"ImageOwnerAlias"`
	SchedulerType   string `position:"Query" name:"SchedulerType"`
}

// RecoverClusterResponse is the response struct for api RecoverCluster
type RecoverClusterResponse struct {
	*responses.BaseResponse
	TaskId    string `json:"TaskId" xml:"TaskId"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateRecoverClusterRequest creates a request to invoke RecoverCluster API
func CreateRecoverClusterRequest() (request *RecoverClusterRequest) {
	request = &RecoverClusterRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("EHPC", "2018-04-12", "RecoverCluster", "ehs", "openAPI")
	request.Method = requests.GET
	return
}

// CreateRecoverClusterResponse creates a response to parse from RecoverCluster response
func CreateRecoverClusterResponse() (response *RecoverClusterResponse) {
	response = &RecoverClusterResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
