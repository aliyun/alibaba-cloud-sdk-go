package selectdb

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

// StartBECluster invokes the selectdb.StartBECluster API synchronously
func (client *Client) StartBECluster(request *StartBEClusterRequest) (response *StartBEClusterResponse, err error) {
	response = CreateStartBEClusterResponse()
	err = client.DoAction(request, response)
	return
}

// StartBEClusterWithChan invokes the selectdb.StartBECluster API asynchronously
func (client *Client) StartBEClusterWithChan(request *StartBEClusterRequest) (<-chan *StartBEClusterResponse, <-chan error) {
	responseChan := make(chan *StartBEClusterResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.StartBECluster(request)
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

// StartBEClusterWithCallback invokes the selectdb.StartBECluster API asynchronously
func (client *Client) StartBEClusterWithCallback(request *StartBEClusterRequest, callback func(response *StartBEClusterResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *StartBEClusterResponse
		var err error
		defer close(result)
		response, err = client.StartBECluster(request)
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

// StartBEClusterRequest is the request struct for api StartBECluster
type StartBEClusterRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	DBClusterId     string           `position:"Query" name:"DBClusterId"`
	DBInstanceId    string           `position:"Query" name:"DBInstanceId"`
}

// StartBEClusterResponse is the response struct for api StartBECluster
type StartBEClusterResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateStartBEClusterRequest creates a request to invoke StartBECluster API
func CreateStartBEClusterRequest() (request *StartBEClusterRequest) {
	request = &StartBEClusterRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("selectdb", "2023-05-22", "StartBECluster", "", "")
	request.Method = requests.POST
	return
}

// CreateStartBEClusterResponse creates a response to parse from StartBECluster response
func CreateStartBEClusterResponse() (response *StartBEClusterResponse) {
	response = &StartBEClusterResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
