package rds

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

// DescribeDBInstanceDetail invokes the rds.DescribeDBInstanceDetail API synchronously
func (client *Client) DescribeDBInstanceDetail(request *DescribeDBInstanceDetailRequest) (response *DescribeDBInstanceDetailResponse, err error) {
	response = CreateDescribeDBInstanceDetailResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDBInstanceDetailWithChan invokes the rds.DescribeDBInstanceDetail API asynchronously
func (client *Client) DescribeDBInstanceDetailWithChan(request *DescribeDBInstanceDetailRequest) (<-chan *DescribeDBInstanceDetailResponse, <-chan error) {
	responseChan := make(chan *DescribeDBInstanceDetailResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDBInstanceDetail(request)
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

// DescribeDBInstanceDetailWithCallback invokes the rds.DescribeDBInstanceDetail API asynchronously
func (client *Client) DescribeDBInstanceDetailWithCallback(request *DescribeDBInstanceDetailRequest, callback func(response *DescribeDBInstanceDetailResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDBInstanceDetailResponse
		var err error
		defer close(result)
		response, err = client.DescribeDBInstanceDetail(request)
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

// DescribeDBInstanceDetailRequest is the request struct for api DescribeDBInstanceDetail
type DescribeDBInstanceDetailRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	ResourceGroupId      string           `position:"Query" name:"ResourceGroupId"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeDBInstanceDetailResponse is the response struct for api DescribeDBInstanceDetail
type DescribeDBInstanceDetailResponse struct {
	*responses.BaseResponse
	ActivationState string `json:"ActivationState" xml:"ActivationState"`
	DBInstanceId    string `json:"DBInstanceId" xml:"DBInstanceId"`
	LicenseType     string `json:"LicenseType" xml:"LicenseType"`
	RequestId       string `json:"RequestId" xml:"RequestId"`
	RegionId        string `json:"RegionId" xml:"RegionId"`
}

// CreateDescribeDBInstanceDetailRequest creates a request to invoke DescribeDBInstanceDetail API
func CreateDescribeDBInstanceDetailRequest() (request *DescribeDBInstanceDetailRequest) {
	request = &DescribeDBInstanceDetailRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "DescribeDBInstanceDetail", "rds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeDBInstanceDetailResponse creates a response to parse from DescribeDBInstanceDetail response
func CreateDescribeDBInstanceDetailResponse() (response *DescribeDBInstanceDetailResponse) {
	response = &DescribeDBInstanceDetailResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
