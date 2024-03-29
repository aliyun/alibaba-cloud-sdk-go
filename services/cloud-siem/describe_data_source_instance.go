package cloud_siem

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

// DescribeDataSourceInstance invokes the cloud_siem.DescribeDataSourceInstance API synchronously
func (client *Client) DescribeDataSourceInstance(request *DescribeDataSourceInstanceRequest) (response *DescribeDataSourceInstanceResponse, err error) {
	response = CreateDescribeDataSourceInstanceResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDataSourceInstanceWithChan invokes the cloud_siem.DescribeDataSourceInstance API asynchronously
func (client *Client) DescribeDataSourceInstanceWithChan(request *DescribeDataSourceInstanceRequest) (<-chan *DescribeDataSourceInstanceResponse, <-chan error) {
	responseChan := make(chan *DescribeDataSourceInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDataSourceInstance(request)
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

// DescribeDataSourceInstanceWithCallback invokes the cloud_siem.DescribeDataSourceInstance API asynchronously
func (client *Client) DescribeDataSourceInstanceWithCallback(request *DescribeDataSourceInstanceRequest, callback func(response *DescribeDataSourceInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDataSourceInstanceResponse
		var err error
		defer close(result)
		response, err = client.DescribeDataSourceInstance(request)
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

// DescribeDataSourceInstanceRequest is the request struct for api DescribeDataSourceInstance
type DescribeDataSourceInstanceRequest struct {
	*requests.RpcRequest
	CloudCode            string `position:"Body" name:"CloudCode"`
	AccountId            string `position:"Body" name:"AccountId"`
	DataSourceInstanceId string `position:"Body" name:"DataSourceInstanceId"`
}

// DescribeDataSourceInstanceResponse is the response struct for api DescribeDataSourceInstance
type DescribeDataSourceInstanceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateDescribeDataSourceInstanceRequest creates a request to invoke DescribeDataSourceInstance API
func CreateDescribeDataSourceInstanceRequest() (request *DescribeDataSourceInstanceRequest) {
	request = &DescribeDataSourceInstanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloud-siem", "2022-06-16", "DescribeDataSourceInstance", "cloud-siem", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeDataSourceInstanceResponse creates a response to parse from DescribeDataSourceInstance response
func CreateDescribeDataSourceInstanceResponse() (response *DescribeDataSourceInstanceResponse) {
	response = &DescribeDataSourceInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
