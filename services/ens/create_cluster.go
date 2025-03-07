package ens

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

// CreateCluster invokes the ens.CreateCluster API synchronously
func (client *Client) CreateCluster(request *CreateClusterRequest) (response *CreateClusterResponse, err error) {
	response = CreateCreateClusterResponse()
	err = client.DoAction(request, response)
	return
}

// CreateClusterWithChan invokes the ens.CreateCluster API asynchronously
func (client *Client) CreateClusterWithChan(request *CreateClusterRequest) (<-chan *CreateClusterResponse, <-chan error) {
	responseChan := make(chan *CreateClusterResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateCluster(request)
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

// CreateClusterWithCallback invokes the ens.CreateCluster API asynchronously
func (client *Client) CreateClusterWithCallback(request *CreateClusterRequest, callback func(response *CreateClusterResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateClusterResponse
		var err error
		defer close(result)
		response, err = client.CreateCluster(request)
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

// CreateClusterRequest is the request struct for api CreateCluster
type CreateClusterRequest struct {
	*requests.RpcRequest
	ClusterVersion string `position:"Query" name:"ClusterVersion"`
	Name           string `position:"Query" name:"Name"`
}

// CreateClusterResponse is the response struct for api CreateCluster
type CreateClusterResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	ClusterId string `json:"ClusterId" xml:"ClusterId"`
}

// CreateCreateClusterRequest creates a request to invoke CreateCluster API
func CreateCreateClusterRequest() (request *CreateClusterRequest) {
	request = &CreateClusterRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ens", "2017-11-10", "CreateCluster", "ens", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateClusterResponse creates a response to parse from CreateCluster response
func CreateCreateClusterResponse() (response *CreateClusterResponse) {
	response = &CreateClusterResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
