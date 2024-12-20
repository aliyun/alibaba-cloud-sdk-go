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

// DescribeCluster invokes the ens.DescribeCluster API synchronously
func (client *Client) DescribeCluster(request *DescribeClusterRequest) (response *DescribeClusterResponse, err error) {
	response = CreateDescribeClusterResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeClusterWithChan invokes the ens.DescribeCluster API asynchronously
func (client *Client) DescribeClusterWithChan(request *DescribeClusterRequest) (<-chan *DescribeClusterResponse, <-chan error) {
	responseChan := make(chan *DescribeClusterResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeCluster(request)
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

// DescribeClusterWithCallback invokes the ens.DescribeCluster API asynchronously
func (client *Client) DescribeClusterWithCallback(request *DescribeClusterRequest, callback func(response *DescribeClusterResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeClusterResponse
		var err error
		defer close(result)
		response, err = client.DescribeCluster(request)
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

// DescribeClusterRequest is the request struct for api DescribeCluster
type DescribeClusterRequest struct {
	*requests.RpcRequest
	ClusterId string `position:"Query" name:"ClusterId"`
}

// DescribeClusterResponse is the response struct for api DescribeCluster
type DescribeClusterResponse struct {
	*responses.BaseResponse
	RequestId string         `json:"RequestId" xml:"RequestId"`
	Clusters  []ClustersItem `json:"Clusters" xml:"Clusters"`
}

// CreateDescribeClusterRequest creates a request to invoke DescribeCluster API
func CreateDescribeClusterRequest() (request *DescribeClusterRequest) {
	request = &DescribeClusterRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ens", "2017-11-10", "DescribeCluster", "ens", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeClusterResponse creates a response to parse from DescribeCluster response
func CreateDescribeClusterResponse() (response *DescribeClusterResponse) {
	response = &DescribeClusterResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
