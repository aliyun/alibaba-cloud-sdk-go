package ddoscoo

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

// DescribeAttackAnalysisMaxQps invokes the ddoscoo.DescribeAttackAnalysisMaxQps API synchronously
func (client *Client) DescribeAttackAnalysisMaxQps(request *DescribeAttackAnalysisMaxQpsRequest) (response *DescribeAttackAnalysisMaxQpsResponse, err error) {
	response = CreateDescribeAttackAnalysisMaxQpsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeAttackAnalysisMaxQpsWithChan invokes the ddoscoo.DescribeAttackAnalysisMaxQps API asynchronously
func (client *Client) DescribeAttackAnalysisMaxQpsWithChan(request *DescribeAttackAnalysisMaxQpsRequest) (<-chan *DescribeAttackAnalysisMaxQpsResponse, <-chan error) {
	responseChan := make(chan *DescribeAttackAnalysisMaxQpsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeAttackAnalysisMaxQps(request)
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

// DescribeAttackAnalysisMaxQpsWithCallback invokes the ddoscoo.DescribeAttackAnalysisMaxQps API asynchronously
func (client *Client) DescribeAttackAnalysisMaxQpsWithCallback(request *DescribeAttackAnalysisMaxQpsRequest, callback func(response *DescribeAttackAnalysisMaxQpsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeAttackAnalysisMaxQpsResponse
		var err error
		defer close(result)
		response, err = client.DescribeAttackAnalysisMaxQps(request)
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

// DescribeAttackAnalysisMaxQpsRequest is the request struct for api DescribeAttackAnalysisMaxQps
type DescribeAttackAnalysisMaxQpsRequest struct {
	*requests.RpcRequest
	EndTime   requests.Integer `position:"Query" name:"EndTime"`
	StartTime requests.Integer `position:"Query" name:"StartTime"`
	SourceIp  string           `position:"Query" name:"SourceIp"`
}

// DescribeAttackAnalysisMaxQpsResponse is the response struct for api DescribeAttackAnalysisMaxQps
type DescribeAttackAnalysisMaxQpsResponse struct {
	*responses.BaseResponse
	Qps       int64  `json:"Qps" xml:"Qps"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDescribeAttackAnalysisMaxQpsRequest creates a request to invoke DescribeAttackAnalysisMaxQps API
func CreateDescribeAttackAnalysisMaxQpsRequest() (request *DescribeAttackAnalysisMaxQpsRequest) {
	request = &DescribeAttackAnalysisMaxQpsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ddoscoo", "2020-01-01", "DescribeAttackAnalysisMaxQps", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeAttackAnalysisMaxQpsResponse creates a response to parse from DescribeAttackAnalysisMaxQps response
func CreateDescribeAttackAnalysisMaxQpsResponse() (response *DescribeAttackAnalysisMaxQpsResponse) {
	response = &DescribeAttackAnalysisMaxQpsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
