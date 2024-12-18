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

// DescribeDBClusterConfigChangeLogs invokes the selectdb.DescribeDBClusterConfigChangeLogs API synchronously
func (client *Client) DescribeDBClusterConfigChangeLogs(request *DescribeDBClusterConfigChangeLogsRequest) (response *DescribeDBClusterConfigChangeLogsResponse, err error) {
	response = CreateDescribeDBClusterConfigChangeLogsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDBClusterConfigChangeLogsWithChan invokes the selectdb.DescribeDBClusterConfigChangeLogs API asynchronously
func (client *Client) DescribeDBClusterConfigChangeLogsWithChan(request *DescribeDBClusterConfigChangeLogsRequest) (<-chan *DescribeDBClusterConfigChangeLogsResponse, <-chan error) {
	responseChan := make(chan *DescribeDBClusterConfigChangeLogsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDBClusterConfigChangeLogs(request)
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

// DescribeDBClusterConfigChangeLogsWithCallback invokes the selectdb.DescribeDBClusterConfigChangeLogs API asynchronously
func (client *Client) DescribeDBClusterConfigChangeLogsWithCallback(request *DescribeDBClusterConfigChangeLogsRequest, callback func(response *DescribeDBClusterConfigChangeLogsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDBClusterConfigChangeLogsResponse
		var err error
		defer close(result)
		response, err = client.DescribeDBClusterConfigChangeLogs(request)
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

// DescribeDBClusterConfigChangeLogsRequest is the request struct for api DescribeDBClusterConfigChangeLogs
type DescribeDBClusterConfigChangeLogsRequest struct {
	*requests.RpcRequest
	StartTime    string `position:"Query" name:"StartTime"`
	ConfigKey    string `position:"Query" name:"ConfigKey"`
	DBInstanceId string `position:"Query" name:"DBInstanceId"`
	DBClusterId  string `position:"Query" name:"DBClusterId"`
	EndTime      string `position:"Query" name:"EndTime"`
}

// DescribeDBClusterConfigChangeLogsResponse is the response struct for api DescribeDBClusterConfigChangeLogs
type DescribeDBClusterConfigChangeLogsResponse struct {
	*responses.BaseResponse
	DynamicCode        string `json:"DynamicCode" xml:"DynamicCode"`
	DynamicMessage     string `json:"DynamicMessage" xml:"DynamicMessage"`
	AccessDeniedDetail string `json:"AccessDeniedDetail" xml:"AccessDeniedDetail"`
	RequestId          string `json:"RequestId" xml:"RequestId"`
	Data               Data   `json:"Data" xml:"Data"`
}

// CreateDescribeDBClusterConfigChangeLogsRequest creates a request to invoke DescribeDBClusterConfigChangeLogs API
func CreateDescribeDBClusterConfigChangeLogsRequest() (request *DescribeDBClusterConfigChangeLogsRequest) {
	request = &DescribeDBClusterConfigChangeLogsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("selectdb", "2023-05-22", "DescribeDBClusterConfigChangeLogs", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeDBClusterConfigChangeLogsResponse creates a response to parse from DescribeDBClusterConfigChangeLogs response
func CreateDescribeDBClusterConfigChangeLogsResponse() (response *DescribeDBClusterConfigChangeLogsResponse) {
	response = &DescribeDBClusterConfigChangeLogsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
